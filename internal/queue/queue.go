package queue

import (
	"notification_dispatch/internal/job"
	"sync"
)

type Status string

const (
	Pending Status = "PENDING"
	Running Status = "RUNNING"
	Done    Status = "DONE"
	Failed  Status = "FAILED"
)

type Queue struct {
	jobs   chan job.Job
	status map[string]Status
	mu     sync.RWMutex
}

func (q *Queue) Enqueue(job job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.status[job.ID] = Pending

	q.jobs <- job
}

func (q *Queue) Dequeue() job.Job {
	return <-q.jobs
}
