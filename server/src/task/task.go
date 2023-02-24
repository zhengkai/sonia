package task

import (
	"project/worker"
)

// Task ...
type Task struct {
	worker []*worker.Worker
	ch     chan *worker.Cmd
}

// NewTask ...
func NewTask(serverList []string) (t *Task) {
	t = &Task{
		ch: make(chan *worker.Cmd),
	}
	for _, host := range serverList {
		t.worker = append(t.worker, worker.NewWorker(host, t.ch))
	}
	return
}

// Do ...
func (t *Task) Do(cmd *worker.Cmd) {
	t.ch <- cmd
}
