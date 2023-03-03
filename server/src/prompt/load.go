package prompt

import (
	"os"
	"project/config"
	"project/util"
	"project/zj"
	"sort"
	"strings"
)

// LoadDir ...
func LoadDir(dir string) (g *Group, err error) {

	rd, err := os.ReadDir(dir)
	if err != nil {
		return
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

	g = &Group{}

	for _, name := range ls {
		file := config.PromptDir + `/` + name
		if strings.HasPrefix(name, `loop_`) {
			loadPrompt(file, &g.PoolLoop)
			continue
		}
		if strings.HasPrefix(name, `rand_`) {
			loadPrompt(file, &g.PoolRand)
			continue
		}
	}
	g.Negative = loadNegative(dir)

	return
}

func loadNegative(dir string) (s string) {

	neg := dir + `/negative.txt`
	if !util.FileExists(neg) {
		zj.J(`negative prompt file not exists, use default`)
		return
	}

	ab, err := os.ReadFile(neg)
	if err != nil {
		zj.W(`load negative prompt file failed:`, neg, err)
		return
	}

	return string(ab)
}
func loadPrompt(file string, kl *[]Keyword) {
	k, err := LoadKeyword(file)
	if err != nil {
		return
	}
	*kl = append(*kl, k)
}
