package worker_pool

import (
	"fmt"
	"time"
)

type WorkerPoolPractice struct {
	jobs	chan int
	results	chan int
}

func NewWorkerPoolPractice() *WorkerPoolPractice{
	return &WorkerPoolPractice{}
}


func (w *WorkerPoolPractice) Start(){
	for i := 1; i <= 100; i++ {
		go w.worker()
	}
}

func (w *WorkerPoolPractice) worker() {
	for j := range w.jobs {
		fmt.Println("started  job", j)
		time.Sleep(time.Second)
		fmt.Println("finished job", j)
		w.results <- j * 2
	}
}


func (w *WorkerPoolPractice) SomeWork(){
	numJobs := 10
	w.jobs = make(chan int, numJobs)
	w.results = make(chan int, numJobs)

	for j := 1; j < numJobs; j++ {
		w.jobs <- j
	}

	for a := 1; a < numJobs; a++ {
		fmt.Println(<-w.results)
	}
}