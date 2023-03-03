package config

// config
var (
	Debug bool
	Dir   string

	Width  = uint32(512)
	Height = uint32(512)

	StaticDir = `/tmp`

	HiRes          = false
	CopyRemoteFile = false

	PromptDir = ``

	Target = `http://127.0.0.1:7860`
)
