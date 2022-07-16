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
	slog.SetupGlobalLogger()
	configmanager.GetGlobalConfig().GlobalLoad()
	slog.SetDisableStandardLogOutput(configmanager.GetGlobalConfig().GetDisableStandardLogOutput())
	static.SetRootDirectory(configmanager.GetGlobalConfig().GetRootDirectory())
	static.SetRedirectSubpaths(configmanager.GetGlobalConfig().GetRedirectSubpaths())
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
