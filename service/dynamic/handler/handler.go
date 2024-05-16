package handler

type HandlerInterface interface {
	Init()
	HandleRequest() error
	GetRequestStruct() interface{}
	GetResponseStruct() interface{}
}
