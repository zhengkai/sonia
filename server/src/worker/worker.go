package worker

import "project/connector"

// Worker ...
type Worker struct {
	con    *connector.Con
	active bool
}

// NewWorker ...
func NewWorker(host, baseDir string, isWindows bool) (w *Worker) {
	w = &Worker{
		con: connector.NewCon(host, baseDir, true),
	}
	go w.background()
	return
}

func (w *Worker) background() {

	w.status()

	// do something
}
