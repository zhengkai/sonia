package task

import (
	"project/worker"
	"time"
)

// Task ...
type Task struct {
	worker []*worker.Worker
}

// Run ...
func (t *Task) Run() {
	time.Sleep(time.Hour)
}
