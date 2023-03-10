package connector

import (
	"bytes"
	"io"
	"net/http"
	"project/metrics"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var fetch = zu.FetchURL

var jsonOpt = protojson.MarshalOptions{
	EmitUnpopulated: true,
}

func postJSON(url string, c proto.Message) (ab []byte, err error) {

	j, err := jsonOpt.Marshal(c)
	if err != nil {
		return
	}

	ab, err = post(url, j)
	if err != nil {
		return
	}
	return
}

func post(url string, content []byte) (ab []byte, err error) {

	metrics.ConnectorReq()
	rsp, err := http.Post(url, `application/json`, bytes.NewBuffer(content))
	metrics.ConnectorRsp()
	if err != nil {
		metrics.ConnectorFail()
		return
	}

	ab, err = io.ReadAll(io.LimitReader(rsp.Body, 1e7))
	if err != nil {
		metrics.ConnectorFail()
		return
	}

	return
}
