package about

import (
	"net/http"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/slog"
)

// ServeAbout
func ServeAbout(r *http.Request, responseContent interface{}) bool {
	res, ok := responseContent.(*globaldefine.VersionInfoPrototype)
	if (!ok) || (res == nil) {
		slog.Errorf("fail to assert version info prototype")
		return false
	}
	res.BaseGoVersion = globaldefine.BaseGoVersion
	res.SoftwareVersion = globaldefine.SoftwareVersion
	res.ProjectHome = globaldefine.ProjectHome
	res.Author = globaldefine.Author
	res.Email = globaldefine.Email
	res.BuildDate = globaldefine.BuildDate
	return true
}
