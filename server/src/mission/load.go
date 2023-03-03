package mission

import (
	"log"
	"os"
	"project/config"
	"project/util"
	"project/zj"
	"sort"
	"strings"
)

var commonNegativePrompt = `(worst quality:2), (low quality:2), (normal quality:2), lowres, normal quality, ((monochrome)), ((grayscale)), skin spots, acnes, skin blemishes, age spot, glan, lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, 3 legs, 3 arms`
var promptLoop []util.Keyword
var promptRand []util.Keyword

func load() {

	rd, err := os.ReadDir(config.PromptDir)
	if err != nil {
		log.Fatal(err)
	}

	var ls []string
	for _, e := range rd {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, `.txt`) {
			continue
		}
		ls = append(ls, name)
	}

	sort.Strings(ls)
	zj.J(`load prompt files:`, ls)

	for _, name := range ls {
		file := config.PromptDir + `/` + name
		if strings.HasPrefix(name, `loop_`) {
			loadPrompt(file, &promptLoop)
			continue
		}
		if strings.HasPrefix(name, `rand_`) {
			loadPrompt(file, &promptRand)
			continue
		}
	}
	loadNegative()
}

func loadNegative() {

	neg := config.PromptDir + `/negative.txt`
	if !util.FileExists(neg) {
		zj.J(`negative prompt file not exists, use default`)
		return
	}

	ab, err := os.ReadFile(neg)
	if err != nil {
		zj.W(`load negative prompt file failed:`, neg, err)
		return
	}

	commonNegativePrompt = string(ab)
}
func loadPrompt(file string, kl *[]util.Keyword) {
	k, err := util.LoadKeyword(file)
	if err != nil {
		return
	}
	*kl = append(*kl, k)
}
