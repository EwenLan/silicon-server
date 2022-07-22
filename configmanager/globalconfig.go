package configmanager

import (
	"encoding/json"

	"github.com/EwenLan/silicon-server/slog"
)

var globalConfigInst *globalConfig

type globalConfigPrototype struct {
	Port                     int
	DisableStandardLogOutput bool
	RootDirectory            string
	LogDirectory             string
	DefaultPage              string
	RedirectSubpaths         []string
}

type globalConfig struct {
	baseConfigure
	configContent       *globalConfigPrototype
	defaultValueLoading bool
}

func (j *globalConfig) GlobalLoad() {
	j.init(defaultGlobalConfigPath)
	j.load()
	j.configContent = &globalConfigPrototype{}
	err := json.Unmarshal(j.fileContent, j.configContent)
	if err != nil {
		slog.Errorf("fail to unmarshal global config file = %s, err = %s", defaultGlobalConfigPath, err)
		j.defaultValueLoading = true
	}
	j.defaultValueLoading = false
}

// GetGlobalConfig
func GetGlobalConfig() *globalConfig {
	if globalConfigInst == nil {
		globalConfigInst = &globalConfig{}
	}
	return globalConfigInst
}

func (j *globalConfig) getConfigBool(varName string, defaultValue bool, value bool) bool {
	if j.defaultValueLoading {
		slog.Debugf("get %s failed, get default value = %t", varName, defaultValue)
		return defaultValue
	}
	slog.Debugf("get %s = %t", varName, value)
	return value
}

func (j *globalConfig) getConfigString(varName string, defaultValue string, value string) string {
	if j.defaultValueLoading {
		slog.Debugf("get %s failed, get default value = %s", varName, defaultValue)
		return defaultValue
	}
	slog.Debugf("get %s = %s", varName, value)
	return value
}

// GetPort
func (j *globalConfig) GetPort() int {
	if j.defaultValueLoading {
		slog.Debugf("get port failed, get default value = %d", defaultGlobalConfig.Port)
		return defaultGlobalConfig.Port
	}
	slog.Debugf("get port = %d", j.configContent.Port)
	return j.configContent.Port
}

// GetDisableStandardLogOutput
func (j *globalConfig) GetDisableStandardLogOutput() bool {
	return j.getConfigBool("standard log output", defaultGlobalConfig.DisableStandardLogOutput, j.configContent.DisableStandardLogOutput)
}

// GetRootDirectory
func (j *globalConfig) GetRootDirectory() string {
	return j.getConfigString("root directory", defaultGlobalConfig.RootDirectory, j.configContent.RootDirectory)
}

// GetLogDirectory
func (j *globalConfig) GetLogDirectory() string {
	return j.getConfigString("log directory", defaultGlobalConfig.LogDirectory, j.configContent.LogDirectory)
}

// GetDefaultPage
func (j *globalConfig) GetDefaultPage() string {
	if (j.defaultValueLoading) || (j.configContent.DefaultPage == "") {
		slog.Debugf("get default page failed, get default value = %s", defaultGlobalConfig.DefaultPage)
		return defaultGlobalConfig.DefaultPage
	}
	slog.Debugf("get default page = %s", j.configContent.DefaultPage)
	return j.configContent.DefaultPage
}

// GetRedirectSubpaths
func (j *globalConfig) GetRedirectSubpaths() []string {
	if j.defaultValueLoading {
		slog.Debugf("get redirect sub-paths failed, get default value = %+v", defaultGlobalConfig.RedirectSubpaths)
		return defaultGlobalConfig.RedirectSubpaths
	}
	slog.Debugf("get redirect sub-paths = %+v", j.configContent.RedirectSubpaths)
	return j.configContent.RedirectSubpaths
}
