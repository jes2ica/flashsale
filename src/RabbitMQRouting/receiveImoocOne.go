package main

import "RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQRouting("exImooc","imooc_one")
	imoocOne.RecieveRouting()
}
