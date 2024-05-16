package service

import (
	"net/http"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic"
	"github.com/EwenLan/silicon-server/service/static"
	"github.com/EwenLan/silicon-server/slog"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	dynamicPrefixLen := len(globaldefine.DynamicPrefix)
	if (len(url) >= dynamicPrefixLen) && (url[:dynamicPrefixLen] == globaldefine.DynamicPrefix) {
		slog.Debugf("path = %s matched with dynamic service", url)
		dynamic.ServeDynamic(w, r)
		return
	}
	slog.Debugf("path = %s static service", url)
	static.ServeStatic(w, r)
}
