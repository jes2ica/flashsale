package main

import (
	"RabbitMQ"
	"fmt"
	)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("Hello imooc")
	fmt.Println("success")
}
