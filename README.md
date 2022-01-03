# errhandle
A golang error handling framework

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## About errhandle:
- An error handling framework built using golang and it is easy to use

## Installation:
```bash
$ go get -u github.com/krishpranav/errhandle
```

## Usage:
```golang
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

```

## Author
- [krishpranav](https://github.com/krishpranav)