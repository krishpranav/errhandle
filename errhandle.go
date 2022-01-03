package errhandle

import (
	"reflect"
)

func NewErrorHandler() *ErrorHandler {
	handler := ErrorHandler{}
	handler.Init()
	return &handler
}

type ErrorHandler struct {
	handlers       map[reflect.Type]func(error, ...interface{}) interface{}
	defaultHandler func(error, ...interface{}) interface{}
}

func (errorHandler *ErrorHandler) Init() {
	errorHandler.handlers = map[reflect.Type](func(error, ...interface{}) interface{}){}
}

func (errorHandler *ErrorHandler) AddHandler(handleFunc func(error, ...interface{}) interface{},
	errTypes ...interface{}) {

	if errTypes == nil {
		errorHandler.defaultHandler = handleFunc
		return
	}
	errType := reflect.TypeOf(errTypes[0])
	errorHandler.handlers[errType] = handleFunc
}

func (errorHandler ErrorHandler) Check(err error, data ...interface{}) interface{} {
	if err == nil {
		return nil
	}
	errType := reflect.TypeOf(err)
	handler, valid := errorHandler.handlers[errType]
	if !valid {
		handler = errorHandler.defaultHandler
	}
	return handler(err, data)
}
