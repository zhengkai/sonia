package build

import "fmt"

// BuildGoVersion ...
var BuildGoVersion string

// BuildTime ...
var BuildTime string

// BuildType ...
var BuildType string

// BuildHost ...
var BuildHost string

// BuildGit ...
var BuildGit string

// DumpBuildInfo ...
func DumpBuildInfo() {
	fmt.Println()
	fmt.Println(BuildGoVersion)
	fmt.Println(BuildTime)
	fmt.Println(BuildType)
	fmt.Println(BuildHost)
	fmt.Println(BuildGit)
	fmt.Println()
}
