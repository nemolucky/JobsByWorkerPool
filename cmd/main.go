package main

import (
	"fmt"

	"github.com/nemolucky/JobsByWorkerPool/internal/domain"
	"github.com/nemolucky/JobsByWorkerPool/internal/producer"
	"github.com/nemolucky/JobsByWorkerPool/internal/worker"
)

func main() {

	jobs := make(chan domain.Job)
	result := make(chan domain.Result)

	jobProducer := producer.NewJobProducer(100)

	workerPool, err := worker.NewPool(10, func(id int) domain.Worker {
		return &worker.JobWorker{ID: id}
	})
	if err != nil {
		fmt.Printf("error create worker pool: %s", err)
		return
	}

	go jobProducer.Generate(jobs)

	go workerPool.Work(jobs, result)

	for res := range result {
		fmt.Printf("Result Original: %d Value: %d\n", res.Original, res.Value)
	}
}
