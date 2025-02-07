package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.NewRabitMQPubSub("imoocPub")
	log.Printf("神爱世人❤️❤️❤️")
	rq.ReceivedPub()
}
