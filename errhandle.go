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

func (errorHandler *ErrorHandler) AddHandler(handleFunc func(error, ...interface{}) interface{},
	errTypes ...interface{}) {

	if errTypes == nil {
		errorHandler.defaultHandler = handleFunc
		return
	}
	errType := reflect.TypeOf(errTypes[0])
	errorHandler.handlers[errType] = handleFunc
}
