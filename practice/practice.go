package practice

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

