package rabbitmq

import (
	"github.com/open-kingfisher/king-utils/common/log"
	"github.com/streadway/amqp"
)

func ProducerPublish(address, exchangeName string, messageBody []byte) error {
	// address: "amqp://saltshaker:saltshaker@localhost:5672/"
	conn, err := amqp.Dial(address)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ:", err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Error("Failed to open a channel:", err)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Error("Failed to declare a exchange:", err)
		return err
	}

	err = ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		})
	if err != nil {
		log.Error("Failed to publish a message:", err)
		return err
	}
	return nil
}
