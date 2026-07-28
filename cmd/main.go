package main

import (
	"fmt"
	"notification_dispatch/internal/job"
	"time"
)

func send(job job.Job) {
	fmt.Println("Sending", job.Type, "notification:", job.Payload)
	time.Sleep(1 * time.Second) // pretend it takes time
	fmt.Println("Done sending job", job.ID)
}

func main() {
	job := job.Job{ID: "1", Type: "email", Payload: "Welcome!"}
	send(job)
}
