package main

import (
	"fmt"

	"github.com/krishpranav/errhandle"
)

func main() {
	myerrHandler := errhandle.NewErrorHandler()

	var SomeData string

	myerrHandler.AddHandler(func(err error, data ...interface{}) interface{} {
		fmt.Println("Generic Error Occured")
		return SomeData
	})

}
