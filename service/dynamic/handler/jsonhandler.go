package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/EwenLan/silicon-server/slog"
)

type JsonHandler struct {
	ServiceHandler HandlerInterface
}

func (v *JsonHandler) innerHandle(reqBody []byte) ([]byte, error) {
	if (len(reqBody) > 0) && (v.ServiceHandler.GetRequestStruct() != nil) {
		slog.Debugf("unmarshal request body %s", string(reqBody))
		err := json.Unmarshal(reqBody, v.ServiceHandler.GetRequestStruct())
		if err != nil {
			slog.Errorf("unmarshal request for body %s failed, err: %s", reqBody, err)
			return nil, err
		}
	}
	err1 := v.ServiceHandler.HandleRequest()
	if err1 != nil {
		slog.Errorf("handle request for body %s failed, err: %s", reqBody, err1)
		return nil, err1
	}
	res, err2 := json.Marshal(v.ServiceHandler.GetResponseStruct())
	if err2 != nil {
		slog.Errorf("marshal response for body %+v failed, err: %s", v.ServiceHandler.GetResponseStruct(), err2)
		return nil, err2
	}
	return res, nil
}

// HttpHandle 处理HTTP请求
func (v *JsonHandler) HttpHandle(w http.ResponseWriter, r *http.Request) {
	reqBuff, err1 := io.ReadAll(r.Body)
	if err1 != nil {
		slog.Errorf("fail to read request for %s", r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err1.Error()))
		return
	}
	res, err := v.innerHandle(reqBuff)
	if err != nil {
		slog.Errorf("inner handle for request %s failed", r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Init 初始化业务处理
func (v *JsonHandler) Init() {
	v.ServiceHandler.Init()
}
