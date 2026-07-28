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

func NewQueue(size int) *Queue {
	return &Queue{
		jobs:   make(chan job.Job, size),
		status: make(map[string]Status),
	}
}

func (q *Queue) Enqueue(job job.Job) {
	q.mu.Lock()
	q.status[job.ID] = Pending
	q.mu.Unlock()

	q.jobs <- job
}

func (q *Queue) Dequeue() job.Job {

	return <-q.jobs
}
func (q *Queue) UpdateStatus(id string, status Status) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.status[id] = status
}

func (q *Queue) GetStatus(id string) (Status, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	status, ok := q.status[id]
	return status, ok
}
