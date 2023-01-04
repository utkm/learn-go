package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// Interface names are the method name with an "er" (convention)
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}


type Incrementer interface {
	Increment() int // returns int
}

type IntCounter int

// Add a method to the custom type (which is the implementation for the Incrementer interface)
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Best Practices
// Use many, small interfaces
// Don't export interfaces for types that will be consumed
// Do export interfaces for types that wil be used by package
// Design functions and methods to receive interfaces whenever possible

// Summary
// Define interface (method signatures inside {}), implement the interface by creating a method on the type
// Implicit: have the methods there that match the type signature for the interface's method
// Empty interface: var i interface{} = 0 (every type in Go implements the empty interface)
// We often pair this with type switches
// Implementing with values vs pointers:
	// Method set of value is all methods with value receivers
	// Method set of pointer is all methods, regardless of receiver type
// Interfaces can be embedded, just like structs