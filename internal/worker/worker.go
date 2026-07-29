package worker

import (
	"notification_dispatch/internal/job"
	"notification_dispatch/internal/queue"
	"notification_dispatch/internal/sender"
)

type Worker struct {
	ID      int
	Queue   *queue.Queue
	Senders map[job.JobType]sender.Sender
}

func NewWorker(
	id int,
	q *queue.Queue,
	senders map[job.JobType]sender.Sender,
) *Worker {
	return &Worker{
		ID:      id,
		Queue:   q,
		Senders: senders,
	}
}

func (w *Worker) Start() {
	for {
		// 1. Take a job from queue
		jb := w.Queue.Dequeue()

		// 2. Mark Running
		w.Queue.UpdateStatus(jb.ID, queue.Running)

		// 3. Find sender for job type
		s, ok := w.Senders[jb.Type]
		if !ok || s == nil {
			w.Queue.UpdateStatus(jb.ID, queue.Failed)
			continue
		}

		// 4. Send
		err := s.Send(jb)

		if err != nil {
			w.Queue.UpdateStatus(jb.ID, queue.Failed)
		} else {
			w.Queue.UpdateStatus(jb.ID, queue.Done)
		}
	}
}
