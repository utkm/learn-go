package main

import (
	"fmt"
	"sync"
)

// A wait group is designed to synchronize multiple Goroutines together
var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	
	// Tell the wait group that we want to synchronize to this Goroutine
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		// Decrements number of groups the wait group is waiting on
		wg.Done()
	}(msg)

	msg = "Goodbye"
	
	// Exit application
	wg.Wait()
}

// Best Practices
// Don't create goroutines in libraries
// When creating a goroutine, know how it will end (avoids subtle memory leaks)
// Check for race conditions at compile time by adding the race flag (go run -race /.../)

// Summary
// Creating goroutines
	// Use go keyword in front of a function call
	// When using anonymous functions, pass data as local variables
// Synchronization
	// Use waitgroups to wait for groups of goroutines to complete
	// Use sync.Mutex and sync.RWMutex to protect data access (only 1 goroutine is manipulating data at once)
// Parallelism
	// By default, Go will use CPU threads equal to available cores
	// More threads can increase performance, but too many can slow it down