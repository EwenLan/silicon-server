package main

import (
	"fmt"
	"net/http"

	"github.com/EwenLan/silicon-server/configmanager"
	"github.com/EwenLan/silicon-server/slog"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	slog.Debugf("Hello world")
	port := configmanager.GetGlobalConfig().GetPort()
	slog.SetStandardLogOutput(configmanager.GetGlobalConfig().GetStandardLogOutput())
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
}
