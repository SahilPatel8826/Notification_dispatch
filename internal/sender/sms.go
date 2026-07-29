package sender

import (
	"fmt"
	"notification_dispatch/internal/job"
	"time"
)

type SmsSender struct{}

func (e *SmsSender) Send(j job.Job) error {
	fmt.Printf("Sending Email | Job ID: %s | Message: %s\n", j.ID, j.Payload)

	// Simulate network delay
	time.Sleep(1 * time.Second)

	fmt.Printf("Email Sent Successfully | Job ID: %s\n", j.ID)

	return nil
}
