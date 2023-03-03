package config

import "os"

func isDir(path string) (b bool) {
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if !info.IsDir() {
		return
	}
	return true
}
