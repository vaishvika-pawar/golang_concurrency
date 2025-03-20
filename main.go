package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ... ?", done)
	go greet("I hope you're liking the course!", done)
	<-done
	<-done
	<-done
	<-done
	// All functions start after the function before is executed completely
	// To run then in parallel we add go keyword before the function call
	// this indicates to call go routines
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// <- this operator sends data to channel
}

// Goroutine behavior
// The idea behind running goroutine is to run it in a non-blocking way

// Channels
// A value that can be used as a communication when working with goroutines
