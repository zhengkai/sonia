package worker

import (
	"time"
)

func (w *Worker) status() {

	first := true

	for {
		// do something

		if !first {
			time.Sleep(1)
		}
		first = false

		rsp, err := w.con.Progress(false)
		if err != nil {
			continue
		}

		w.active = rsp.Active
	}
}
