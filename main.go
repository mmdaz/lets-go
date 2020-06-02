package main

import (
	c "lets_go/concurrency_practice"
)

func main() {
	//cmd.Main()
	// learn_kafka.StartReading()
	// learn_kafka.Produce()
	//fast_http_example.StartServer()
	//log.Logger.Info("main started...")

	var users []int
	finished := make(chan bool)
	for index := 0; index < 80; index++ {
		users = append(users, index)
	}
	c.PracticeWaitGroups(users, finished)
	<- finished
}
