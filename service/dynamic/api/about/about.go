package about

import (
	"net/http"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/handler"
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

type AboutHandler struct {
	response globaldefine.VersionInfoPrototype
}

func (i *AboutHandler) HandleRequest() error {
	i.response.BaseGoVersion = globaldefine.BaseGoVersion
	i.response.SoftwareVersion = globaldefine.SoftwareVersion
	i.response.ProjectHome = globaldefine.ProjectHome
	i.response.Author = globaldefine.Author
	i.response.Email = globaldefine.Email
	i.response.BuildDate = globaldefine.BuildDate
	return nil
}

func (i *AboutHandler) GetRequestStruct() interface{} {
	return nil
}

func (i *AboutHandler) GetResponseStruct() interface{} {
	return &i.response
}

var About = handler.JsonHandler{
	ServiceHandler: &AboutHandler{},
}
