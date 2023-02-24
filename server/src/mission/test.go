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
const commonNegativePrompt = `(worst quality:2), (low quality:2), (normal quality:2), lowres, normal quality, ((monochrome)), ((grayscale)), skin spots, acnes, skin blemishes, age spot, glan, lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, 3 legs, 3 arms`

// Test for dev
func Test() {

	list := []string{`http://10.0.32.2:7860`}

	t := task.NewTask(list)

	for i, v := range anmials {
		for j, w := range place {

			zj.J(i, j, v, w)

			seed := 1

			prompt := fmt.Sprintf(`%s, ((%s)), %s on %s`, commonPrompt, v, v, w)

			file := fmt.Sprintf(`output/abc/%d_%d.png`, i, j)

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
}
