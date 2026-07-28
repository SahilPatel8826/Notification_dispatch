package job

type Job struct {
	ID      string
	Type    string // "email", "sms", "push"
	Payload string // e.g. the message content
}
