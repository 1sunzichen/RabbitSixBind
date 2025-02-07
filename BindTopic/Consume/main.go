package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.NewRabitMQTopic("imoocTopic", "imooc.*.one")
	log.Printf("神爱世人❤️❤️❤️")
	rq.ReceivedTopic()
}
