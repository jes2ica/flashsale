package main

import "RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.*.two")
	imoocOne.RecieveTopic()
}
