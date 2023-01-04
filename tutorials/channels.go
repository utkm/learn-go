package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// Channels are designed to synchronize data transmission between multiple goroutines
	// Note that channels are strongly typed
	ch := make(chan int, 50)
	wg.Add(2)
	
	// This is the receiving goroutine
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	// This is the sending goroutine 
	go func(ch chan<- int) {
		ch <- 21
		ch <- 19
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

// Summary
// Basics
	// Create a channel with the make keyword
	// Can send a message into a channel (ch <- val)
	// Can receive a message from a channel (val := <-ch)
	// Can have multiple senders and receivers
// Restricting data flow
	// Channels can be cast into send-only or receive-only versions
// Buffered channels
	// Channels block sender side till receiver is available
	// Block receiver side till message is available
	// Can decouple sender and receiver with buffered channels
	// Use buffered channels when send and receiver have asymmetric loading
// For range loops with channels
	// The first parameter you receive is the value itself, not the index
	// Used to monitor channel and process messages as they arrive
	// Loop exits when channel is closed
// Select statements
	// Work kind of like switch statements but only in the context of channels
	// Allows goroutine to monitor several channels at once
	// Blocks if all channels block
	// If multiple channels receive value simultaneously, behaviour is undefined (so ordering doesn't matter)
	// 