package project

import (
	"project/build"
	"project/config"
	"project/mission"
	"project/zj"
	"time"
)

// Start ...
func Start() {

	build.DumpBuildInfo()

	zj.Init()

	mission.Test()

	time.Sleep(time.Hour)
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
