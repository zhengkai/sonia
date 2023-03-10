package connector

import (
	"encoding/json"
	"fmt"
	"project/config"
	"project/pb"
	"project/util"
	"project/zj"
	"strings"
	"time"

	jp "github.com/buger/jsonparser"
)

/*
{
    "data": [
        [
            {
                "name": "D:\\ai\\stable-diffusion-webui\\outputs\\txt2img-images\\2023-02-23\\00021-3719897399.png",
                "data": null,
                "is_file": true
            }
        ],
        "{\"prompt\": \"\", \"all_prompts\": [\"\"], \"negative_prompt\": \"\", \"all_negative_prompts\": [\"\"], \"seed\": 3719897399, \"all_seeds\": [3719897399], \"subseed\": 110177673, \"all_subseeds\": [110177673], \"subseed_strength\": 0, \"width\": 512, \"height\": 512, \"sampler_name\": \"Euler a\", \"cfg_scale\": 7, \"steps\": 20, \"batch_size\": 1, \"restore_faces\": false, \"face_restoration_model\": null, \"sd_model_hash\": \"d8691b4d16\", \"seed_resize_from_w\": 0, \"seed_resize_from_h\": 0, \"denoising_strength\": null, \"extra_generation_params\": {}, \"index_of_first_image\": 0, \"infotexts\": [\"Steps: 20, Sampler: Euler a, CFG scale: 7, Seed: 3719897399, Size: 512x512, Model hash: d8691b4d16, Model: deliberate_v11\"], \"styles\": [], \"job_timestamp\": \"20230223172606\", \"clip_skip\": 1, \"is_using_inpainting_conditioning\": false}",
        "<p>Steps: 20, Sampler: Euler a, CFG scale: 7, Seed: 3719897399, Size: 512x512, Model hash: d8691b4d16, Model: deliberate_v11</p>",
        "<p></p><div class='performance'><p class='time'>Time taken: <wbr>1.93s</p><p class='vram'>Torch active/reserved: 3160/3274 MiB, <wbr>Sys VRAM: 5477/10240 MiB (53.49%)</p></div>"
    ],
    "is_generating": false,
    "duration": 1.9338173866271973,
    "average_duration": 4.410197890323142
}
*/

type predictReq struct {
	Data        any    `json:"data"`
	FnIndex     int    `json:"fn_index"`
	SessionHash string `json:"session_hash"`
}

// PredictRsp ...
type PredictRsp []byte

func (c *Con) predict(fn int, data any) (out PredictRsp, err error) {

	req := &predictReq{
		Data:        data,
		FnIndex:     fn,
		SessionHash: c.hash,
	}
	ab, err := json.Marshal(req)
	if err != nil {
		return
	}
	zj.J(`in`, len(ab))
	util.WriteFile(`predict-in.json`, ab)

	ab, err = post(c.url(`/run/predict/`), ab)
	if err != nil {
		return
	}

	out = ab

	util.WriteFile(`predict-out.json`, out)
	return
}

// Predict ...
func (c *Con) Predict(p *pb.Predict) (out PredictRsp, err error) {

	predictMakeupDefaultValue(p)

	d := []any{
		c.id,
		p.Prompt,
		p.NegativePrompt,
		[]int{},
		p.Steps,
		p.SamplerName,
		p.RestoreFaces,
		false,
		1,
		1,
		7,
		-1,
		-1,
		0,
		0,
		0,
		false,
		p.Height,
		p.Width,
		config.HiRes,
		0.7,
		2,
		"Latent",
		0,
		0,
		0,
		[]int{},
		"None",
		false,
		false,
		"positive",
		"comma",
		0,
		false,
		false,
		"",
		"Seed",
		"",
		"Nothing",
		"",
		"Nothing",
		"",
		true,
		false,
		false,
		false,
		0,
		predictFile(p.Seed, c.isWindows, c.baseDir),
		"{\"prompt\": \"flying pig\", \"all_prompts\": [\"flying pig\"], \"negative_prompt\": \"\", \"all_negative_prompts\": [\"\"], \"seed\": 3420343523, \"all_seeds\": [3420343523], \"subseed\": 2009181532, \"all_subseeds\": [2009181532], \"subseed_strength\": 0, \"width\": 512, \"height\": 512, \"sampler_name\": \"Euler a\", \"cfg_scale\": 7, \"steps\": 20, \"batch_size\": 1, \"restore_faces\": false, \"face_restoration_model\": null, \"sd_model_hash\": \"d8691b4d16\", \"seed_resize_from_w\": 0, \"seed_resize_from_h\": 0, \"denoising_strength\": null, \"extra_generation_params\": {}, \"index_of_first_image\": 0, \"infotexts\": [\"flying pig\\nSteps: 20, Sampler: Euler a, CFG scale: 7, Seed: 3420343523, Size: 512x512, Model hash: d8691b4d16, Model: deliberate_v11\"], \"styles\": [], \"job_timestamp\": \"20230223174441\", \"clip_skip\": 1, \"is_using_inpainting_conditioning\": false}",
		"<p></p>",
		"<p></p>",
	}
	return c.predict(77, d)
}

// Resize ...
func (c *Con) Resize(w, h uint32, rate float64) (out PredictRsp, err error) {

	d := []any{
		true,
		w,
		h,
		rate,
		0,
		0,
	}
	return c.predict(64, d)
}

func predictFile(seed uint32, isWindows bool, baseDir string) []*pb.PredictFile {

	date := time.Now().Format(`20060102`)
	time := time.Now().Format(`150405`)

	file := fmt.Sprintf(`%s/%s/%s-%d.png`, baseDir, date, time, seed)
	if isWindows {
		file = strings.ReplaceAll(file, `/`, `\`)
	}

	return []*pb.PredictFile{
		{
			Data:   `file=` + file,
			Name:   file,
			IsFile: true,
		},
	}
}

func predictMakeupDefaultValue(p *pb.Predict) {

	if p.SamplerName == `` {
		p.SamplerName = `Euler a`
	}
	if p.Steps == 0 {
		p.Steps = 20
	}
	if p.Width == 0 {
		p.Width = config.Width
	}
	if p.Height == 0 {
		p.Height = config.Height
	}
}

// GetFile ...
func (p PredictRsp) GetFile() (file string, err error) {
	file, err = jp.GetString(p, `data`, `[0]`, `[0]`, `name`)
	return
}
