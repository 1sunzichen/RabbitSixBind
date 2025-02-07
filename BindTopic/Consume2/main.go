package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.NewRabitMQTopic("imoocTopic", "#")
	log.Printf("神爱世人❤️❤️❤️2")
	rq.ReceivedTopic()
}
