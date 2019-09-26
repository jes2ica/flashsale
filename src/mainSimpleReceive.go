package main

import "RabbitMQ"

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
