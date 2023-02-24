package connector

import (
	"fmt"
	"math/rand"
)

// Con ...
type Con struct {
	id        string
	server    string
	hash      string
	isWindows bool
	baseDir   string
}

// NewCon ...
func NewCon(server string) (c *Con) {
	return &Con{
		server: server,
		id:     fmt.Sprintf(`task(%s)`, randomID(15)),
		hash:   randomID(10),
	}
}

// GetServer ...
func (c *Con) GetServer() string {
	return c.server
}

func randomID(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	result[0] = chars[rand.Intn(len(chars))]
	for i := 1; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func (c *Con) url(path string) string {
	return c.server + path
}
