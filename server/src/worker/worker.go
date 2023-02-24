package worker

import (
	"project/connector"
	"project/pb"
	"project/zj"
	"time"
)

// Worker ...
type Worker struct {
	con    *connector.Con
	active bool
}

// NewWorker ...
func NewWorker(host string) (w *Worker) {
	w = &Worker{
		con: connector.NewCon(host),
	}
	go w.background()
	return
}

func (w *Worker) background() {

	for {
		err := w.con.BaseDir()
		if err == nil {
			break
		}
		zj.W(w.con.GetHost(), `fetch fail`, err, `, retry after 5s...`)
		time.Sleep(time.Second * 5)
	}

	go w.status()

	w.predict()
	// do something
}

func (w *Worker) predict() {
	w.con.Predict(&pb.Predict{
		Prompt: `a flying pig on the moon`,
		Height: 1600,
		Width:  1600,
	})
}
