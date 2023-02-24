package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"project/config"
	"project/zj"
	"strings"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/proto"
)

// DownloadFunc ...
type DownloadFunc func(url string) (ab []byte, err error)

// IsURL ...
func IsURL(s string) bool {
	return strings.HasPrefix(s, `https://`) || strings.HasPrefix(s, `http://`)
}

// StaticFile ...
func StaticFile(file string) string {
	file = strings.TrimPrefix(file, config.StaticDir+`/`)
	return fmt.Sprintf(`%s/%s`, config.StaticDir, file)
}

// ReadFile ...
func ReadFile(file string) (ab []byte, err error) {
	ab, err = os.ReadFile(StaticFile(file))
	zj.Watch(&err)
	return
}

// FileExists ...
func FileExists(file string) bool {
	return zu.FileExists(StaticFile(file))
}

// SaveData ...
func SaveData(name string, p proto.Message) (err error) {

	defer zj.Watch(&err)

	ab, err := proto.Marshal(p)
	zj.J(name, err)
	if err == nil {
		WriteFile(name+`.pb`, ab)
	}

	ab, err = json.Marshal(p)
	if err == nil {
		WriteFile(name+`.json`, ab)
	}

	return
}

// WriteFile ...
func WriteFile(file string, ab []byte) (err error) {

	defer zj.Watch(&err)

	file = StaticFile(file)

	f, err := os.CreateTemp(config.StaticDir+`/tmp`, `wr-*`)
	if err != nil {
		return
	}
	tmpName := f.Name()
	_, err = f.Write(ab)
	if err != nil {
		return
	}

	/*
		brErr := zu.Brotli(tmpName)
		gzErr := zu.Gzip(tmpName)
		if len(ab) > 1000 {
			defer func() {
				if err != nil {
					return
				}
				if brErr == nil {
					os.Rename(tmpName+`.br`, file+`.br`)
				}
				if gzErr == nil {
					os.Rename(tmpName+`.gz`, file+`.gz`)
				}
			}()
		}
	*/

	err = os.Chmod(tmpName, 0644)
	if err != nil {
		return
	}

	dir := path.Dir(file)
	if !zu.FileExists(dir) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}

	err = os.Rename(tmpName, file)
	if err != nil {
		return
	}

	return
}

// MarkSync ...
func MarkSync() {
	zj.J(`mark sync`)
	WriteFile(`sync.txt`, []byte{65})
}

// Download ...
func Download(url, file string, fn DownloadFunc) (err error) {

	if url == `` || file == `` {
		err = errors.New(`no string`)
		return
	}

	file = StaticFile(file)
	if zu.FileExists(file) {
		return
	}

	ab, err := fn(url)
	if err != nil {
		zj.J(`fetch fail`, url)
		return
	}

	if len(ab) == 0 {
		zj.J(`fetch fail, 0 size`, url)
		return
	}

	f, err := os.CreateTemp(config.StaticDir+`/tmp`, `download-*`)
	if err != nil {
		return
	}
	_, err = f.Write(ab)
	if err != nil {
		return
	}
	f.Close()

	os.Chmod(f.Name(), 0644)
	os.Rename(f.Name(), file)
	return
}
