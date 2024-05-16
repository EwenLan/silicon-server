package dynamic

import (
	"net/http"

	"github.com/EwenLan/silicon-server/slog"
)

// ServeDynamic
func ServeDynamic(w http.ResponseWriter, r *http.Request) {
	guider := &guiderType{}
	guider.init(r.Method, r.URL.Path)
	handle := rootRoutineNode.searchRoutineNode(guider)
	if handle == nil {
		slog.Errorf("path = %+v is unable to find", guider.steps)
		http.NotFound(w, r)
		return
	}
	handle.Init()
	handle.HttpHandle(w, r)
	slog.Debugf("found handle for path = %+v", guider.steps)
}
