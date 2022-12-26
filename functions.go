package main

import (
	"fmt"
)

// Takes no parameters, andd returns no values
func main() {
	greeting := "Hello"
	name := "Jo"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s)

	for i := 0; i < 5; i++ {
		func(i int) { // should set up a parameter so that changes in the outerscope aren't reflected in the inner
			fmt.Println(i)
		}(i) // invoke the function
	}

	var f func() = func() {
		fmt.Println("Hello func")
	}
	f()

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// Takes a pointer (much more efficient since it's copied otherwise)
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// Variadic Parameters (can only have one and it has to be at the end)
func sum(values ...int) *int {
	fmt.Println(values)
	// result is declared on the execution stack of this function
	// when the function executes, the execution stack is destroyed which could be problemtic
	// but in Go, when it recognizes that you're returning a value generated on the local stack, it promotes it to the heap (shared memory)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

// method is a function that is executing in a known context (any type)
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Summary
// Same naming conventions apply (if first letter is capitalized, it's exported)
// Variadic parameters are received as a slice
// If you have a single return value, just list the type
// If there are multiple return values, list types surrounded by parentheses
// Can also used named return values, and returns this used the return keyword on its own
// Can return addresses of local variables
// Anonymous functions:
// Functions don't have names if they are immediately invoked or assigned to a variable
// or passed as an argument to a function
// Functions are types
// can assignmen them to variables or use as arguments and return values in functions
// type signature is like function signature, with no parameter names
// var f func(string, string, int) (int, error)
// Methods:
// Function that executes in the context of a type (any custom type)
// func (g greeter) greet() {}
// g is the value receiver so we get a copy of that greeter object and that's going to be passed into the greet method
// Can also use pointer receivers (usually a much more efficient operation, since you get to avoid copying)
