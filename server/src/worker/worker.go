package worker

import (
	"project/config"
	"project/connector"
	"project/pb"
	"project/zj"
	"time"
)

// Worker ...
type Worker struct {
	con    *connector.Con
	active bool
	ch     chan *Cmd
	prevW  uint32
	prevH  uint32
}

// NewWorker ...
func NewWorker(server string, ch chan *Cmd) (w *Worker) {
	w = &Worker{
		con: connector.NewCon(server),
		ch:  ch,
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
		zj.W(w.con.GetServer(), `fetch fail`, err, `, retry after 5s...`)
		time.Sleep(time.Second * 5)
	}

	// go w.status()
	w.waitActive()

	for {
		cmd, ok := <-w.ch
		if !ok {
			break
		}
		w.work(cmd)
	}
}

func (w *Worker) work(c *Cmd) (err error) {

	defer zj.Watch(&err)

	cp := c.Predict
	if config.HiRes {
		if cp.Width != w.prevW || cp.Height != w.prevH {
			w.prevW = cp.Width
			w.prevH = cp.Height
			ab, err := w.con.Resize(cp.Width, cp.Height, 2)
			zj.J(`resize`, err, string(ab))
		}
	}

	p, err := w.con.Predict(cp)
	if err != nil {
		return
	}
	return

	f, err := p.GetFile()
	if err != nil {
		return
	}

	_, err = w.con.Download(f, c.FileName)
	if err != nil {
		return
	}
	return
}

func (w *Worker) waitActive() {
	for {
		rsp, _ := w.con.Progress(false)
		if rsp.GetActive() {
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func (w *Worker) testPredict() {
	w.con.Predict(&pb.Predict{
		Prompt: `a flying pig on the moon`,
	})
}
