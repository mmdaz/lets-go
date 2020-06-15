package main

import "lets_go/worker_pool_practice/worker_pool"

func main() {
	workerPool := worker_pool.NewWorkerPoolPractice()
	workerPool.Start()
	for i := 0; i < 50000; i++ {
		workerPool.SomeWork()
	}
}




