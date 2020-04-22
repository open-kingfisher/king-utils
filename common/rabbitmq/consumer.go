package rabbitmq

import (
	"encoding/json"
	"github.com/open-kingfisher/king-utils/common"
	"github.com/open-kingfisher/king-utils/common/log"
	"github.com/open-kingfisher/king-utils/kit"
	"github.com/streadway/amqp"
	"time"
)

type MQHandler interface {
	HandleMessage(m *amqp.Delivery) error
}

type UpdateKubeConfig struct{}

func (u *UpdateKubeConfig) HandleMessage(m *amqp.Delivery) error {
	if len(m.Body) == 0 {
		return nil
	}
	cluster := common.ClusterDB{}
	if err := json.Unmarshal(m.Body, &cluster); err != nil {
		log.Error(err)
		return err
	}
	kubeConfig := common.KubeConfigPath + cluster.Id
	log.Infof("cluster id: %s, create kubeconfig file", cluster.Id)
	if err := kit.CreateConfig(cluster, kubeConfig); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

type Consumer struct {
	Address            string
	ExchangeName       string
	Handler            MQHandler
	connNotifyClose    chan *amqp.Error
	channelNotifyClose chan *amqp.Error
}

func (c *Consumer) Run() {
	conn, err := amqp.Dial(c.Address)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ:", err)
		time.Sleep(5 * time.Second)
		// 重连
		c.Run()
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Error("Failed to open a channel:", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		c.ExchangeName, // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Error("Failed to declare a exchange:", err)
		return
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Error("Failed to declare a queue:", err)
		return
	}

	err = ch.QueueBind(
		q.Name,         // queue name
		"",             // routing key
		c.ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to bind a queue:", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Error("Failed to register a consumer:", err)
		return
	}
	// 保存NotifyClose和NotifyClose用于连接关闭和channel关闭后进行重连
	c.connNotifyClose = conn.NotifyClose(make(chan *amqp.Error))
	c.channelNotifyClose = ch.NotifyClose(make(chan *amqp.Error))
	go c.ReConnect()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			func(h MQHandler) {
				if err := h.HandleMessage(&d); err != nil {
					return
				}
			}(c.Handler)
		}
	}()

	<-forever
}

func (c *Consumer) ReConnect() {
	for {
		select {
		case err := <-c.connNotifyClose:
			if err != nil {
				log.Error("Connection NotifyClose:", err)
				c.Run()
			}
		case err := <-c.channelNotifyClose:
			if err != nil {
				log.Error("Channel NotifyClose:", err)
				c.Run()
			}
		}
	}
}
