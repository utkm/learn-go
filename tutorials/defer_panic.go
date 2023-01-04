package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// // Go executes functions passed into it after the function finishes its final statement but before it returns
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// // Deferred statements are executed in LIFO order
	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	// Get a response and a possible error
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Defer done working with web request
	// Allows you to associate openening and closing the resource right next to each other
	defer res.Body.Close() 
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)


	// Panic statements happen after defer statements are evaluated
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("panic attack")
	// fmt.Println("end")
}

// Summary
// Defer is used to delay execution of a statement until the function exits
// It's useful to group open and close functions together 
// Defer statements run in LIFO order (pretty intuitive)
// Arguments are evaluated at the time the defer is executed, not at the time of function exiting
// Panic occurs when the application cannot continue at all
// The function will stop executing but the deferred functions will still fire
// If nothing handles the panic, the program will exit
// Can recover from panics using the recover keyword
// It's best used inside deferred functions
// The current function will not attempt to continue, but higher functions in the call stack will
