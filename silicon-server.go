package main

import (
	"fmt"
	"net/http"

	"github.com/EwenLan/silicon-server/configmanager"
	"github.com/EwenLan/silicon-server/service"
	"github.com/EwenLan/silicon-server/service/dynamic"
	"github.com/EwenLan/silicon-server/service/static"
	"github.com/EwenLan/silicon-server/slog"
)

func initalSetup() {
	// 日志初始化
	slog.SetupGlobalLogger()
	// 加载全局配置
	configmanager.GetGlobalConfig().GlobalLoad()
	// 设置日志输出
	slog.SetDisableStandardLogOutput(configmanager.GetGlobalConfig().GetDisableStandardLogOutput())
	// 设置根目录
	static.SetRootDirectory(configmanager.GetGlobalConfig().GetRootDirectory())
	// 设置重定向路径
	static.SetRedirectSubpaths(configmanager.GetGlobalConfig().GetRedirectSubpaths())
	// 初始化路径树
	dynamic.InitRootRoutineNode()
}

func main() {
	initalSetup()
	port := configmanager.GetGlobalConfig().GetPort()
	http.HandleFunc("/", service.Serve)
	slog.Debugf("http server start to listen = http://localhost:%d", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		slog.Errorf("fail to start http server, err = %s", err)
		return
	}
	slog.Debugf("program exit")
}
