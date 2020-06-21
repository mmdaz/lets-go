package worker_pool

import (
	"testing"
)

func BenchmarkWorkerPool(b *testing.B) {
	workerPool := NewWorkerPoolPractice()
	workerPool.Start()
	for i := 0; i < b.N; i++ {
		workerPool.SomeWork()
	}
}
