package handler

type HandlerInterface interface {
	HandleRequest() error
	GetRequestStruct() interface{}
	GetResponseStruct() interface{}
}
