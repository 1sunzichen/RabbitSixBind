package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.RabbitMQSimple("imooc")
	rq.PublishSimple("Welcome to HM ❤️❤️❤️")
	log.Printf("施比受更有福")
}
