package connector

import "project/pb"

// Test for dev
func Test() {

	o := &pb.ProgressReq{}
	postJSON(``, o)
}
