package project

import (
	"project/build"
	"project/config"
	"project/task"
	"project/zj"
	"time"
)

// Start ...
func Start() {

	build.DumpBuildInfo()

	zj.Init()

	task.Test()

	time.Sleep(time.Hour)
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
