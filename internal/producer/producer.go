package producer

import (
	"fmt"

	"github.com/nemolucky/JobsByWorkerPool/internal/domain"
)

type JobProducer struct {
	Quantity int
}

func NewJobProducer(quantity int) *JobProducer {
	return &JobProducer{
		Quantity: quantity,
	}
}

func (jp *JobProducer) Generate(c chan<- domain.Job) {

	defer close(c)

	if jp.Quantity <= 0 {
		fmt.Println("quantity must not be 0 or negative")
		return
	}

	for i := 1; i <= jp.Quantity; i++ {

		createdJob := domain.Job{
			ID:    i,
			Value: (i + 10) % 7,
		}

		c <- createdJob
	}
}
