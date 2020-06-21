package worker_pool

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPoolPractice struct {
	jobs	chan waitStruct
}

type waitStruct struct {
	wait *sync.WaitGroup
	value int
}

func NewWorkerPoolPractice() *WorkerPoolPractice{
	return &WorkerPoolPractice{
		jobs: make(chan waitStruct),
	}
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
		j.wait.Done()
	}
}


func (w *WorkerPoolPractice) SomeWork(){
	numJobs := 200
	var wait sync.WaitGroup
	wait.Add(numJobs)
	go func() {
		for j := 0; j < numJobs; j++ {
			w.jobs <- waitStruct{wait: &wait, value:j}
		}
	}()
	wait.Wait()
}