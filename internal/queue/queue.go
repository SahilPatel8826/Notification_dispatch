package queue

import "notification_dispatch/internal/job"

type Queue struct {
	jobs chan job.Job
}

func (q *Queue) Enqueue(job Job) {
	q.jobs <- job
}

func (q *Queue) Dequeue() Job {
	return <-q.jobs
}
