package main

import (
	"fmt"
	"log"
	RabbitMQ "mqpro/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rq := RabbitMQ.NewRabitMQRouting("imoocRouting", "r1")
	rq2 := RabbitMQ.NewRabitMQRouting("imoocRouting", "r2")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "=>r1")
			rq.PublishRouting("Welcome to HM ❤️❤️❤️" + strconv.Itoa(i))
		} else {
			fmt.Println(i, "=>r2")
			rq2.PublishRouting("Welcome to HM2 ❤️❤️❤️" + strconv.Itoa(i))
		}
		time.Sleep(time.Second)

	}
	log.Printf("施比受更有福-route模式")
}
