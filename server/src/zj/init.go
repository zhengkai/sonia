package zj

import (
	"path/filepath"
	"project/config"

	"github.com/zhengkai/zog"
)

// Init ...
func Init() {

	mainCfg := zog.NewConfig()
	mainCfg.Caller = zog.CallerLong

	infoCfg := mainCfg.Clone()
	infoCfg.Color = zog.ColorInfo
	infoCfg.LinePrefix = `[IO] `

	debugCfg := mainCfg.Clone()
	debugCfg.Color = zog.ColorLight
	debugCfg.LinePrefix = `[Debug] `

	errCfg := zog.NewErrConfig()
	errCfg.Color = zog.ColorWarn
	errCfg.LinePrefix = `[Error] `

	baseLog.CDefault = mainCfg
	baseLog.CDebug = debugCfg
	baseLog.CInfo = infoCfg
	baseLog.CError = errCfg
	baseLog.CWarn = errCfg
	baseLog.CFatal = errCfg

	baseLog.SetDirPrefix(filepath.Dir(zog.GetSourceFileDir()))

	// 生产环境走 docker，不写本地文件
	if !config.Prod {

		mainFile, _ := zog.NewFile(config.Dir+`/log/default.txt`, false)
		infoFile, _ := zog.NewFile(config.Dir+`/log/io.txt`, false)
		errFile, _ := zog.NewFile(config.Dir+`/log/err.txt`, true)

		mainCfg.Output = append(mainCfg.Output, mainFile)
		infoCfg.Output = append(infoCfg.Output, infoFile)
		errCfg.Output = append(errCfg.Output, mainFile, errFile)
	}
}
