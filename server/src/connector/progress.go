package connector

import (
	"project/pb"
	"project/zj"

	"google.golang.org/protobuf/encoding/protojson"
)

// Progress ...
func (c *Con) Progress(isPreview bool) (rsp *pb.ProgressRsp, err error) {

	m := &pb.ProgressReq{
		IdTask: c.id,
	}
	if !isPreview {
		m.IdLivePreview = -1
	}

	url := c.url(`/internal/progress`)
	ab, err := postJSON(url, m)
	if err != nil {
		return
	}

	rsp = &pb.ProgressRsp{}

	err = protojson.Unmarshal(ab, rsp)
	if err != nil {
		zj.J(string(ab))
		zj.J(url)
		zj.W(err)
		return
	}
	return
}
