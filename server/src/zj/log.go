package zj

import "github.com/zhengkai/zog"

var baseLog = &zog.Logger{}

// J log
var J = baseLog.Println

// F log printf
var F = baseLog.Printf

// D debug log
var D = baseLog.Debugln

// DF debug printf
var DF = baseLog.Debugf

// W warn log
var W = baseLog.Warningln

// WF warn log
var WF = baseLog.Warningf

// IO ...
var IO = baseLog.Infoln

// IOF ...
var IOF = baseLog.Infof

// Watch ...
var Watch = baseLog.WatchStack

// N log nothing
func N(x ...interface{}) {
}
