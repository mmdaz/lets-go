package main

import (
	"lets_go/concurrency_practice"
	"lets_go/interface_practice"
)

func main() {
	//cmd.Main()
	// learn_kafka.StartReading()
	// learn_kafka.Produce()
	//fast_http_example.StartServer()
	//log.Logger.Info("main started...")

	var users []int
	finished := make(chan bool)
	for index := 0; index < 50000; index++ {
		users = append(users, index)
	}

	interPractice := interface_practice.NewInterfacePractice()
	waitGroupPractice := concurrency_practice.NewWaitGroupPractice(interPractice)
	waitGroupPractice.RunPracticeWaitGroups(users, finished)

	<-finished
}
