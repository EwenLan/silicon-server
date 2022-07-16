package about

import (
	"encoding/json"
	"net/http"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/slog"
)

type versionInfoPrototype struct {
	BaseGoVersion   string
	SoftwareVersion string
	ProjectHome     string
	Author          string
	Email           string
	BuildDate       string
}

var versionInfo = versionInfoPrototype{
	BaseGoVersion:   globaldefine.BaseGoVersion,
	SoftwareVersion: globaldefine.SoftwareVersion,
	ProjectHome:     globaldefine.ProjectHome,
	Author:          globaldefine.Author,
	Email:           globaldefine.Email,
	BuildDate:       globaldefine.BuildDate,
}

// ServeAbout
func ServeAbout(w http.ResponseWriter, r *http.Request) {
	info, err := json.Marshal(versionInfo)
	if err != nil {
		slog.Errorf("fail to marshal version info, err = %s", err)
		http.NotFound(w, r)
		return
	}
	w.Write(info)
	w.Header().Set("Content-Type", "applicatin/json")
	// w.WriteHeader(200)
}
