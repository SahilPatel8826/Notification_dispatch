package worker

import (
	"notification_dispatch/internal/job"
	"notification_dispatch/internal/queue"
	"notification_dispatch/internal/sender"
)

type Pool struct {
	workers []*Worker
}

func NewPool(
	numWorkers int,
	q *queue.Queue,
	senders map[job.JobType]sender.Sender,
) *Pool {

	pool := &Pool{}

	for i := 1; i <= numWorkers; i++ {
		w := NewWorker(i, q, senders)
		pool.workers = append(pool.workers, w)
	}

	return pool
}

func (p *Pool) Start() {
	for _, worker := range p.workers {
		go worker.Start()
	}
}
