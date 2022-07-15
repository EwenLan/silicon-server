package configmanager

import (
	"encoding/json"

	"github.com/EwenLan/silicon-server/slog"
)

var globalConfigInst *globalConfig

var defaultGlobalConfig globalConfigPrototype = globalConfigPrototype{
	Port:              8080,
	StandardLogOutput: true,
}

type globalConfigPrototype struct {
	Port              int
	StandardLogOutput bool
}

type globalConfig struct {
	baseConfigure
	configContent       *globalConfigPrototype
	defaultValueLoading bool
}

func (j *globalConfig) globalLoad() {
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

// GetPort
func (j *globalConfig) GetPort() int {
	if (j.defaultValueLoading) || (j.configContent.Port == 0) {
		slog.Debugf("get configured port failed, get default port = %d", defaultGlobalConfig.Port)
		return defaultGlobalConfig.Port
	}
	slog.Debugf("get configured port = %d", j.configContent.Port)
	return j.configContent.Port
}

// GetStandardLogOutput
func (j *globalConfig) GetStandardLogOutput() bool {
	if j.defaultValueLoading {
		slog.Debugf("get standard log output failed, get default option = %t", defaultGlobalConfig.StandardLogOutput)
		return defaultGlobalConfig.StandardLogOutput
	}
	slog.Debugf("get standard log output = %t", j.configContent.StandardLogOutput)
	return j.configContent.StandardLogOutput
}

func init() {
	globalConfigInst = &globalConfig{}
	globalConfigInst.globalLoad()
}

// GetGlobalConfig
func GetGlobalConfig() *globalConfig {
	return globalConfigInst
}
