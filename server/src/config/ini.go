package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func runIni() {

	cfg, err := ini.Load(Dir + `/sonia.ini`)
	if err != nil {
		return
	}

	sec := cfg.Section(``)

	Debug = sec.Key(`debug`).MustBool()
	PromptDir = sec.Key(`prompt_dir`).MustString(``)
	if PromptDir != `` && !isDir(PromptDir) {
		fmt.Println(`Invalid prompt_dir:`, PromptDir)
		PromptDir = ``
	}

	w := sec.Key(`width`).MustUint()
	if w >= 256 && w <= 4096 {
		Width = uint32(w)
	}

	h := sec.Key(`height`).MustUint()
	if h >= 256 && h <= 4096 {
		Height = uint32(h)
	}

	Target = sec.Key(`target`).MustString(Target)
}
