package worker

import (
	"project/zj"
	"time"
)

func (w *Worker) status() {

	first := true

	for {
		// do something

		if !first {
			time.Sleep(time.Second * 3)
		}
		first = false

		rsp, err := w.con.Progress(false)
		if err != nil {
			continue
		}

		if rsp.LivePreview != `` {
			zj.W(w.con.GetServer(), `live preview:`, len(rsp.LivePreview))
		}

		w.active = rsp.Active
		zj.J(`active`, w.active)
	}
}
