package worker

import (
	"encoding/base64"
	"errors"
	"regexp"
)

var pBase64Head = regexp.MustCompile(`^data:([^;]+);base64,`)

func readInlineBase64(s string) (mime string, ab []byte, err error) {

	if len(s) < 50 {
		err = errors.New(`too short`)
		return
	}

	head := s[0:50]
	if !pBase64Head.MatchString(head) {
		err = errors.New(`not base64`)
		return
	}

	mime = pBase64Head.FindStringSubmatch(head)[1]

	idx := len(mime) + 5 + 8
	s = s[idx:]

	ab, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		err = errors.New(`base64 decode error`)
		return
	}

	return
}
