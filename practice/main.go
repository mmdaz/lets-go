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

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {

	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	go printCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("End of main")
}
