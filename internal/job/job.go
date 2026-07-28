package job

type Job struct {
	ID      string
	Type    string // "email", "sms", "push"
	Payload string // e.g. the message content
}

func (j *Job) NewJob(id, jobType, payload string) {
	j.ID = id
	j.Type = jobType
	j.Payload = payload
}
