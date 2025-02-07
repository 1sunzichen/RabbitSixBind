package RabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://guest:guest@127.0.0.1:5672/test"

type RabbitMQ struct {
	Conn     *amqp.Connection
	Mqurl    string
	Key      string
	Exchange string
	Qname    string
	Channel  *amqp.Channel
}

func NewRabitMQ(qName string, exChange string, Key string) *RabbitMQ {
	var err error
	mq := &RabbitMQ{
		Qname:    qName,
		Exchange: exChange,
		Key:      Key,
		Mqurl:    MQURL,
	}
	mq.Conn, err = amqp.Dial(mq.Mqurl)
	mq.failOnErr(err, "获取队列失败")
	mq.Channel, err = mq.Conn.Channel()
	mq.failOnErr(err, "获取管道失败")
	return mq

}

func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.Conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}

func RabbitMQSimple(qName string) *RabbitMQ {
	return NewRabitMQ(qName, "", "")
}

func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.Channel.QueueDeclare(r.Qname, false, false, false, false, nil)
	r.failOnErr(err, "声明队列失败")
	r.Channel.Publish(r.Exchange, r.Qname, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}

func (r *RabbitMQ) ConsumeSimple() {
	_, err := r.Channel.QueueDeclare(r.Qname, false, false, false, false, nil)
	r.failOnErr(err, "声明队列失败")
	message, err := r.Channel.Consume(r.Qname, "", false, false, false, false, nil)
	r.failOnErr(err, "消费者声明失败")
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("获取消息:%s", d.Body)
		}
	}()
	log.Printf("Exit:CTRL+C")
	<-forever

}
