package mission

import (
	"fmt"
	"math/rand"
	"project/pb"
	"project/task"
	"project/util"
	"project/worker"
	"project/zj"
	"time"
)

// Test for dev
func Test() {

	load()

	list := []string{`http://127.0.0.1:7860`}

	t := task.NewTask(list)

	idx := uint32(rand.Int())

	for {

		prompt := ``
		for _, v := range promptLoop {
			_, s := v.Index(idx)
			prompt += s
		}
		for _, v := range promptRand {
			_, s := v.Rand()
			prompt += s
		}

		idx++

		zj.J(prompt)
		time.Sleep(time.Second / 5)

		seed := 1

		file := fmt.Sprintf(`output/abc/%d_%d.png`, 1, idx)

		if util.FileExists(file) {
			zj.J(`skip`)
			continue
		}

		file = util.StaticFile(file)

		cmd := &worker.Cmd{
			Predict: &pb.Predict{
				Prompt: prompt,
				Seed:   uint32(seed),
			},
			FileName: file,
		}

		t.Do(cmd)
	}
}
