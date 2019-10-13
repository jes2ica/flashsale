package main

import "RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("exImoocTopic","#")
	imoocOne.RecieveTopic()
}
