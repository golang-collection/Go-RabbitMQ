package main

import "Go-RabbitMQ/tools"

/**
* @Author: super
* @Date: 2020-08-13 08:12
* @Description:
**/

func main() {
	testOne := tools.NewRabbitMQRouting("testRouting", "testTwo")
	testOne.ReceiveRouting()
}