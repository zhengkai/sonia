package prompt

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
)

// ErrKeywordEmpty ...
var ErrKeywordEmpty = errors.New(`keyword is empty`)

// Keyword ...
type Keyword []string

// Index of keyword
func (k Keyword) Index(i uint32) (idx uint32, s string) {
	if len(k) == 0 {
		return
	}
	idx = i % uint32(len(k))
	return idx, k[idx] + `,`
}

// Rand of keyword
func (k Keyword) Rand() (idx uint32, s string) {
	if len(k) == 0 {
		return
	}
	idx = uint32(rand.Intn(len(k)))
	return idx, k[idx] + `,`
}

// LoadKeyword ...
func LoadKeyword(file string) (k Keyword, err error) {

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") || strings.HasPrefix(s, "//") {
			continue
		}
		s = strings.TrimSpace(s)
		s = strings.TrimRight(s, `,`)
		if s == `` {
			continue
		}
		k = append(k, s)
	}
	if len(k) == 0 {
		err = ErrKeywordEmpty
	}
	return
}
