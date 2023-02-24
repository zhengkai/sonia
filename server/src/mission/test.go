package mission

import (
	"fmt"
	"project/pb"
	"project/task"
	"project/util"
	"project/worker"
	"project/zj"
)

const commonPrompt = `masterpiece,best quality,official art,extremely detailed CG unity 8k wallpaper,ultra high res, (photorealistic:1.4)`

// Test for dev
func Test() {

	list := []string{`http://10.0.32.2:7860`}

	t := task.NewTask(list)
	zj.J(t)

	// t.Do()

	for i, v := range anmials {
		for j, w := range place {
			fmt.Println(v, w)

			prompt := fmt.Sprintf(`%s, %s on %s`, commonPrompt, v, w)

			file := fmt.Sprintf(`output/%d_%d.png`, i, j)
			file = util.StaticFile(file)

			cmd := &worker.Cmd{
				Predict: &pb.Predict{
					Prompt: prompt,
				},
				FileName: file,
			}

			t.Do(cmd)
		}
	}
}
