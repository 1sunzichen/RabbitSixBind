package main

import (
	"fmt"
	"log"
	RabbitMQ "mqpro/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rq := RabbitMQ.NewRabitMQTopic("imoocTopic", "imooc.topic.one")
	rq2 := RabbitMQ.NewRabitMQTopic("imoocTopic", "r2")
	for i := 0; i <= 100; i++ {

		rq.PublishTopic("Welcome to HM ❤️❤️❤️" + strconv.Itoa(i))

		rq2.PublishTopic("Welcome to HM2 ❤️❤️❤️" + strconv.Itoa(i))
		fmt.Println(i)
		time.Sleep(time.Second)

	}
	log.Printf("施比受更有福-route模式")
}
