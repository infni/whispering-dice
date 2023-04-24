package main

import "fmt"

// the logging package requires something that adheres to an interface.
// a single function doesn't match to an interface
// wrap the simple function fmt.Println() in a struct so it conforms to the interface.

type customLogger struct{}

func (_ customLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}
