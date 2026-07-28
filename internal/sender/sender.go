package sender

import "notification_dispatch/internal/job"

type Sender interface {
	Send(job.Job) error
}
