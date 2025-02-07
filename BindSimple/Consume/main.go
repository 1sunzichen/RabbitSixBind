package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.RabbitMQSimple("imooc")
	log.Printf("神爱世人❤️❤️❤️")
	rq.ConsumeSimple()
}
