package main

import (
	"fmt"
	"log"
	RabbitMQ "mqpro/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rq := RabbitMQ.RabbitMQSimple("imooc")
	for i := 0; i <= 100; i++ {

		rq.PublishSimple("Welcome to HM ❤️❤️❤️" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
	log.Printf("施比受更有福")
}
