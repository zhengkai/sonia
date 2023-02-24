package worker

import "project/pb"

// Cmd ...
type Cmd struct {
	Predict  *pb.Predict
	FileName string
}
