package concurrency_practice

import (
	"fmt"
	"runtime"
)

func schedule() {

	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again.")
	runtime.Gosched()
}
