package task

import "project/worker"

// Test for dev
func Test() {

	t := &Task{
		worker: []*worker.Worker{
			worker.NewWorker(`http://10.0.32.2:7860`, `D:\ai\stable-diffusion-webui\outputs\txt2img-images\`, true),
		},
	}

	t.Run()
}