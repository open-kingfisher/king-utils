package config

import (
	"flag"
	"os"
	"strings"
)

var (
	Listen      string
	DBURL       string
	HarborURL   string
	RabbitMQURL string
	isHelp      bool
	Mode        string
)

func init() {
	cmd := flag.NewFlagSet("", flag.ContinueOnError)
	cmd.StringVar(&Listen, "listen", "0.0.0.0:8080", "Host Listen")
	cmd.StringVar(&Mode, "mode", "release", "Set Mode, Options: [debug|release|test]")
	cmd.StringVar(&DBURL, "dbURL", "root:password@tcp(10.10.10.10:3306)/kingfisher", "DB URL")
	cmd.StringVar(&HarborURL, "harborURL", "admin:passowrd@registry.kingfihser.com", "Harbor URL")
	cmd.StringVar(&RabbitMQURL, "rabbitMQURL", "amqp://kingfisher:kingfisher@localhost:5672/", "RabbitMQ URL")
	cmd.BoolVar(&isHelp, "help", false, "Print this help")
	cmd.Parse(args(cmd))
	if isHelp {
		cmd.PrintDefaults()
		os.Exit(1)
	}
}

func args(f *flag.FlagSet) (args []string) {
	left := strings.Join(os.Args, " ")
	for _, v := range os.Args {
		key := strings.TrimLeft(strings.Split(v, "=")[0], "-")
		if f.Lookup(key) != nil {
			args = append(args, v)
			left = strings.Replace(left, " "+v, "", -1)
		}
	}
	os.Args = strings.Split(left, " ")
	return args
}
