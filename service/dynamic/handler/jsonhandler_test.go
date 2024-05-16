package handler

import (
	"encoding/json"
	"testing"
)

type requestStruct struct {
	Req string
}

type responseStruct struct {
	Res string
}

type innerHander1 struct {
	request  requestStruct
	response responseStruct
}

func (i *innerHander1) HandleRequest() error {
	i.response.Res = i.request.Req
	return nil
}

func (i *innerHander1) GetRequestStruct() interface{} {
	return &i.request
}

func (i *innerHander1) GetResponseStruct() interface{} {
	return &i.response
}

func (i *innerHander1) Init() {
	i.request = requestStruct{}
	i.response = responseStruct{}
}

func TestJsonHander(t *testing.T) {
	t.Run("check response", func(t *testing.T) {
		handle1 := JsonHandler{ServiceHandler: new(innerHander1)}
		requestStr, _ := json.Marshal(&requestStruct{Req: "hello world"})
		res, _ := handle1.innerHandle(requestStr)
		resStruct := responseStruct{}
		json.Unmarshal(res, &resStruct)
		if resStruct.Res != "hello world" {
			t.Errorf("response is %s", res)
		}
	})
}
