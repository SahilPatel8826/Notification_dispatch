package queue

import "notification_dispatch/internal/job"

func Enqueue(jobs chan Job, job job.Job) {
	jobs <- job
}
func Dequeue(jobs chan Job) Job {
	job := <-jobs
	return job
}
