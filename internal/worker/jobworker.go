package worker

import (
	"fmt"

	"github.com/nemolucky/JobsByWorkerPool/internal/domain"
)

type JobWorker struct {
	ID int
}

func (w *JobWorker) Do(job domain.Job) domain.Result {
	return domain.Result{
		ID:       job.ID,
		Original: job.Value,
		Value:    job.Value * job.Value,
	}
}

func (w JobWorker) About() {
	fmt.Printf("JobWorker ID %d\n", w.ID)
}
