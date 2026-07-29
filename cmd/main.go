package main

import (
	"fmt"
	"time"

	"notification_dispatch/internal/job"
	"notification_dispatch/internal/queue"
	"notification_dispatch/internal/sender"
	"notification_dispatch/internal/worker"
)

func main() {

	// Create Queue
	q := queue.NewQueue(10)

	// Register all senders
	senders := map[job.JobType]sender.Sender{
		job.Email: &sender.EmailSender{},
		job.SMS:   &sender.SmsSender{},
		job.Push:  &sender.PushSender{},
	}

	// Create Worker Pool
	pool := worker.NewPool(3, q, senders)

	// Start Workers
	pool.Start()

	// Create Jobs
	jobs := []job.Job{
		job.NewJob("1", job.Email, "Welcome Email"),
		job.NewJob("2", job.SMS, "OTP : 123456"),
		job.NewJob("3", job.Push, "New Notification"),
		job.NewJob("4", job.Email, "Verify Email"),
		job.NewJob("5", job.SMS, "Transaction Alert"),
		job.NewJob("6", job.Push, "Friend Request"),
	}

	// Submit Jobs
	for _, j := range jobs {
		fmt.Println("Submitting Job:", j.ID)
		q.Enqueue(j)
	}

	// Wait for workers to finish
	time.Sleep(10 * time.Second)

	// Print Job Status
	fmt.Println("\nFinal Status")

	for _, j := range jobs {
		status, _ := q.GetStatus(j.ID)
		fmt.Printf("Job %s -> %s\n", j.ID, status)
	}
}
