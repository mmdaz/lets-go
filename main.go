package main

import (
	"lets_go/concurrency_practice"
	"lets_go/interface_practice/modules"
)

func main() {
	//cmd.Main()
	// learn_kafka.StartReading()
	// learn_kafka.Produce()
	//fast_http_example.StartServer()
	//log.Logger.Info("main started...")
	interPractice := modules.NewMainModule()
	waitGroupPractice := concurrency_practice.NewWaitGroupPractice(interPractice)


	var users []int
	finished := make(chan bool)
	for index := 0; index < 50000; index++ {
		users = append(users, index)
	}

	var divided [][]int
	for i := 0; i < len(users); i += 80 {
		end := i + 80
		if end > len(users) {
			end = len(users)
		}
		divided = append(divided, users[i:end])
	}

	for _, chunkedUsers := range divided {
		waitGroupPractice.RunPracticeWaitGroups(chunkedUsers, finished)
	}



	<-finished
}
