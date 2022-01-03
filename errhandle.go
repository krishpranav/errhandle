package errhandle

import "reflect"

type ErrorHandler struct {
	handlers       map[reflect.Type]func(error, ...interface{}) interface{}
	defaultHandler func(error, ...interface{})
}

func NewErrorHandler() *ErrorHandler {
	handler := ErrorHandler{}
	handler.Init()
	return &handler
}

func (errorHandler *ErrorHandler) Init() {
	errorHandler.handlers = map[reflect.Type](func(error, ...interface{}) interface{}){}
}
