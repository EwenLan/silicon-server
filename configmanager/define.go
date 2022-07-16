package configmanager

const defaultGlobalConfigPath = "GlobalConfig.json"

var defaultGlobalConfig globalConfigPrototype = globalConfigPrototype{
	Port:                     8080,
	DisableStandardLogOutput: false,
	RootDirectory:            "www",
	LogDirectory:             "log",
	DefaultPage:              "index.html",
}
