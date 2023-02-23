package connector

import "project/pb"

func (c *con) Progress() {

	m := &pb.Progress{}

	postJSON(`http://localhost:8080/progress`, m)
}
