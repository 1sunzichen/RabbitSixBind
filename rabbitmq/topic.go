package RabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

// 订阅模式的 声明  生产的消息，可被 消费者 共享消费
// 传入交换机 参数

// 订阅端方法 声明交换机 发布消息
// 消费端方法 声明交换机 声明队列 声明bind关系 开始消费
func NewRabitMQTopic(exchangename string, routingKey string) *RabbitMQ {
	return NewRabitMQ("", exchangename, routingKey)
}

// 生产端
func (r *RabbitMQ) PublishTopic(message string) {
	err := r.Channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "交换机声明失败")
	err = r.Channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "订阅模式 发布失败")
}

// 消费端
func (r *RabbitMQ) ReceivedTopic() {
	err := r.Channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "交换机声明失败")
	q, err := r.Channel.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "声明队列失败")
	err = r.Channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式QueueBind失败")
	message, err := r.Channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费失败")
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("message %s", d.Body)
		}
	}()
	log.Printf("EXIT:CTRL+C")
	<-forever
}
