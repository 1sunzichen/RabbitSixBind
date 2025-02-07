package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.NewRabitMQRouting("imoocRouting", "r1")
	log.Printf("神爱世人❤️❤️❤️")
	rq.ReceivedRouting()
}
