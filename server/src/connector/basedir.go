package connector

import (
	"errors"
	"regexp"
)

var regexpJS = regexp.MustCompile(`<script type="text/javascript" src="file=(.*?)script.js?`)

// BaseDir ...
func (c *Con) BaseDir() (err error) {
	s, err := baseDir(c.url(``))
	if err != nil {
		return
	}
	c.baseDir = s + `outputs/txt2img-image/`

	c.isWindows = c.baseDir[1] == ':'
	return
}

func baseDir(url string) (dir string, err error) {

	ab, err := fetch(url)
	if err != nil {
		return
	}

	m := regexpJS.FindSubmatch(ab)
	if len(m) < 2 {
		err = errors.New(`baseDir not found`)
		return
	}

	dir = string(m[1])
	if len(dir) < 5 {
		err = errors.New(`baseDir syntax error`)
		return
	}

	return
}
