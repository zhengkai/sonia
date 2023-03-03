package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {

	Dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	list := map[string]*string{
		`SONIA_DIR`: &StaticDir,
	}
	for k, v := range list {
		s := os.Getenv(k)
		if len(s) > 1 {
			*v = s
		}
	}

	runIni()

	fmt.Println(`debug:`, Debug)

	s := PromptDir
	if s == `` {
		s = `<empty>`
	}
	fmt.Println(`prompt_dir:`, PromptDir)
}
