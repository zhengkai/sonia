package util

// Keyword ...
type Keyword []string

// Index of keyword
func (k Keyword) Index(i uint32) (idx uint32, s string) {
	if len(k) == 0 {
		return
	}
	idx = i % uint32(len(k))
	return idx, k[idx]
}
