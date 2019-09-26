package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url format: amqp:username:password@host:port/vhost
const MQURL = "amqp://yijiez:12345@127.0.0.1:5672/yijiez"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
	Exchange string
	Key string
	Mqurl string
}

// create flashsale instance
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName:queueName, Exchange:exchange, Key:key, Mqurl:MQURL}
}

// disconnect
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// error handling
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "failed when creating connection")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "failed when getting channel")
	return rabbitMQ
}

func (r *RabbitMQ) PublishSimple(message string) {
	// 1. Apply for the queue, create if queue doesn't exist.
	// Guarantee that the queue exists & message can be delivered to the queue
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// whether to persist or not
		false,
		// when the last consumer disconnect, whether to delete
		false,
		// if true, only the creator can see
		false,
		// whether to block
		false,
		// extra attributes
		nil)

	if err != nil {
		fmt.Println(err)
	}
	// 2. Send the message to the queue
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// Send the message back to the producer if cannot find the queue (based on exchange, type)
		false,
		// If the queue doesn't have consumer, send the message back to the producer.
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	// 1. Apply for the queue.
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// whether to persist or not
		false,
		// when the last consumer disconnect, whether to delete
		false,
		// if true, only the creator can see
		false,
		// if false, block
		false,
		// extra attributes
		nil)

	if err != nil {
		fmt.Println(err)
	}
	// 2. Receive message
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		// whether to ack automatically when message has been consumed
		true,
		false,
		// if true, cannot pass the mesasge to the consumers in the same connection
		false,
		false,
		nil)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit, press ctrl+c")
	<-forever
}


