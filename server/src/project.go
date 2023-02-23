package project

import (
	"project/build"
	"project/config"
	"project/connector"
	"project/zj"
	"time"
)

// Start ...
func Start() {

	build.DumpBuildInfo()

	zj.Init()

	connector.Test()

	time.Sleep(time.Hour)
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
