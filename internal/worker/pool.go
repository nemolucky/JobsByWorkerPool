package worker

import (
	"fmt"
	"sync"

	"github.com/nemolucky/JobsByWorkerPool/internal/domain"
)

type WorkerPool struct {
	Workers []domain.Worker
}

func NewPool(quantity int, factory func(id int) domain.Worker) (*WorkerPool, error) {

	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}

	workers := make([]domain.Worker, 0, quantity)

	for i := 1; i <= quantity; i++ {
		workers = append(workers, factory(i))
	}

	return &WorkerPool{
		Workers: workers,
	}, nil
}

func (wp *WorkerPool) Work(jobs <-chan domain.Job, result chan<- domain.Result) {

	defer close(result)

	wg := sync.WaitGroup{}

	for _, worker := range wp.Workers {
		wg.Go(func() {
			for job := range jobs {
				result <- worker.Do(job)
			}
		})
	}

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}
