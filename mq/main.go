package main

import (
	"log"
	"time"
	"took/mq/mq"
)

func nowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	var (
		conf = mq.Conf{
			User: "guest",
			Pwd:  "guest",
			Addr: "127.0.0.1",
			Port: "5672",
		}

		exchangeName = "user.register.direct"
		queueName    = "user.register.queue"
		keyName      = "user.register.event"
	)

	if err := mq.Init(conf); err != nil {
		log.Fatalf(" mq init err: %v", err)
	}

	ch := mq.NewChannel()
	if err := ch.ExchangeDeclare(exchangeName, "direct"); err != nil {
		log.Fatalf("create exchange err: %v", err)
	}
	if err := ch.QueueDeclare(queueName); err != nil {
		log.Fatalf("create queue err: %v", err)
	}
	if err := ch.QueueBind(queueName, keyName, exchangeName); err != nil {
		log.Fatalf("bind queue err: %v", err)
	}
	select {}
}
