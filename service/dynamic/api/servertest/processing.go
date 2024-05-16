package servertest

import (
	"strings"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/handler"
	"github.com/EwenLan/silicon-server/slog"
)

type ProcessorReusingTest struct {
	request  globaldefine.ProcessorReusingTestRequest
	response globaldefine.ProcessorReusingTestResponse
}

func (i *ProcessorReusingTest) HandleRequest() error {
	slog.Errorf("get request body %+v", i.request)
	if len(i.request.MessageA) > 0 {
		i.response.ResponseA = strings.ToUpper(i.request.MessageA)
	}
	if len(i.request.MessageB) > 0 {
		i.response.ResponseB = strings.ToUpper(i.request.MessageB)
	}
	if len(i.request.MessageC) > 0 {
		i.response.ResponseC = strings.ToUpper(i.request.MessageC)
	}
	return nil
}

func (i *ProcessorReusingTest) GetRequestStruct() interface{} {
	return &i.request
}

func (i *ProcessorReusingTest) GetResponseStruct() interface{} {
	return &i.response
}

func (i *ProcessorReusingTest) Init() {
	i.request = globaldefine.ProcessorReusingTestRequest{}
	i.response = globaldefine.ProcessorReusingTestResponse{}
}

var ProcessorReusingTestImp = handler.JsonHandler{
	ServiceHandler: &ProcessorReusingTest{},
}
