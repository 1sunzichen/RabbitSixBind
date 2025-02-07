package main

import (
	"log"
	RabbitMQ "mqpro/rabbitmq"
)

func main() {
	rq := RabbitMQ.NewRabitMQRouting("imoocRouting", "r2")
	log.Printf("神爱世人❤️❤️❤️2")
	rq.ReceivedRouting()
}
