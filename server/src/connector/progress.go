package connector

import (
	"project/pb"
	"project/zj"

	"github.com/zhengkai/zu"
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

	ab, err := postJSON(c.url(`/internal/progress`), m)
	if err != nil {
		return
	}

	rsp = &pb.ProgressRsp{}

	err = protojson.Unmarshal(ab, rsp)
	if err != nil {
		zj.W(err)
		return
	}

	zj.J(string(ab))
	zj.J(zu.JSON(rsp))
	return
}
