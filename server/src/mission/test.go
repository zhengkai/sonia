package mission

import (
	"fmt"
	"math/rand"
	"project/config"
	"project/prompt"
	"project/task"
	"project/util"
	"project/worker"
	"project/zj"
	"time"
)

// Test for dev
func Test() {

	pg, err := prompt.LoadDir(config.PromptDir)
	if err != nil {
		zj.W(`load prompt failed`, err)
		return
	}

	list := []string{`http://127.0.0.1:7860`}

	t := task.NewTask(list)

	idx := uint32(rand.Int())

	for p := range pg.Infinity() {

		zj.J(p.Prompt)
		time.Sleep(time.Second / 5)

		// seed := 1

		file := fmt.Sprintf(`output/abc/%d_%d.png`, 1, idx)

		if util.FileExists(file) {
			zj.J(`skip`)
			continue
		}

		file = util.StaticFile(file)

		cmd := &worker.Cmd{
			Predict:  p,
			FileName: file,
		}

		t.Do(cmd)
	}
}
