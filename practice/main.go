package main

import (
	"fmt"
	"time"
)

func count() {
	for index := 0; index < 5; index++ {
		fmt.Println(index)
		time.Sleep(time.Millisecond * 1000)
	}

}

func main() {

	go count()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5000)

}
