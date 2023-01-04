package main

import (
	"fmt"
	// "runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS(1) application runs single-threadded. A truly concurrent application with no parallelism at all
	// Useful in situations where there's a lot of data synchronization and you need to be careful to avoid race conditions that parallelism can cause

	for i := 0; i < 10; i++ {
		wg.Add(2)

		// Locks are executing before each Goroutine executes
		// Locking the Mutexes in a single context (but now concurrency and parallelism are basically destroyed)
		// All of the Mutexes are forcing the data to be synchronized and are forced it to be run in a single threaded way
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	// Trying to protect the counter variable from concurrent reading and writing
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// Goroutines are racing against each other -- there's no synchronization
// We need to synchronize them together -- can use Mutex (a lock that the application will honour)