package servertest

import (
	"fmt"
	"strconv"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/handler"
	"github.com/EwenLan/silicon-server/slog"
)

type Calculator struct {
	request  globaldefine.CalculatingRequest
	response globaldefine.CalculatingResponse
}

type CalculateFunc = func(float64, float64) float64

var calculateMap = map[string]CalculateFunc{
	"plus":  func(a float64, b float64) float64 { return a + b },
	"minus": func(a float64, b float64) float64 { return a - b },
	"times": func(a float64, b float64) float64 { return a * b },
	"divid": func(a float64, b float64) float64 { return a / b },
}

func (i *Calculator) HandleRequest() error {
	numA, err1 := strconv.ParseFloat(i.request.NumA, 64)
	if err1 != nil {
		slog.Errorf("fail to parse numA: %s", i.request.NumA)
		return err1
	}
	numB, err2 := strconv.ParseFloat(i.request.NumB, 64)
	if err2 != nil {
		slog.Errorf("fail to parse numB: %s", i.request.NumB)
		return err1
	}
	f, ok := calculateMap[i.request.Op]
	if (f == nil) || (!ok) {
		slog.Errorf("unknown operation: %s", i.request.Op)
		return fmt.Errorf("unknown operation: %s", i.request.Op)
	}
	i.response.Ans = strconv.FormatFloat(f(numA, numB), 'f', -1, 64)
	return nil
}

func (i *Calculator) GetRequestStruct() interface{} {
	return &i.request
}

func (i *Calculator) GetResponseStruct() interface{} {
	return &i.response
}

func (i *Calculator) Init() {
	i.request = globaldefine.CalculatingRequest{}
	i.response = globaldefine.CalculatingResponse{}
}

var CalculatorImp = handler.JsonHandler{
	ServiceHandler: &Calculator{},
}
