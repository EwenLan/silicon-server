package jsonhandler

import (
	"encoding/json"
	"net/http"

	"github.com/EwenLan/silicon-server/slog"
)

type jsonHandleFuncType func(*http.Request, interface{}) bool

// JsonHandle
type JsonHandle struct {
	ResponseContent interface{}
	JsonHandleFunc  jsonHandleFuncType
}

// HttpHandle
func (j *JsonHandle) HttpHandle(w http.ResponseWriter, r *http.Request) {
	if j.JsonHandleFunc == nil {
		slog.Errorf("json handle func is nil")
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	if !j.JsonHandleFunc(r, j.ResponseContent) {
		slog.Errorf("fail to get response content")
		w.WriteHeader(http.StatusBadRequest)
	}
	buff, err := json.Marshal(j.ResponseContent)
	if err != nil {
		slog.Errorf("fail to marshal json, err = %s", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Write(buff)
}
