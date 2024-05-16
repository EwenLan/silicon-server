package about

import (
	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/handler"
)

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

func (i *AboutHandler) Init() {
	i.response = globaldefine.VersionInfoPrototype{}
}

var About = handler.JsonHandler{
	ServiceHandler: &AboutHandler{},
}
