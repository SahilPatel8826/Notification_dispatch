package job

type JobType string

const (
	Email JobType = "email"
	SMS   JobType = "sms"
	Push  JobType = "push"
)

type Job struct {
	ID      string
	Type    JobType
	Payload string
}

func NewJob(id string, jobType JobType, payload string) Job {
	return Job{
		ID:      id,
		Type:    jobType,
		Payload: payload,
	}
}
