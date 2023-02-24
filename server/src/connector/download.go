package connector

import (
	"net/url"
	"project/util"
)

// Download ...
func (c *Con) Download(remote, local string) (size int, err error) {

	p := url.Values{}
	p.Add(`file`, remote)

	ab, err := fetch(c.server + `/` + p.Encode())
	if err != nil {
		return
	}

	size = len(ab)

	util.WriteFile(local, ab)
	return
}
