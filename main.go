package main

import (
	"fmt"
	"time"
)

func main() {
	dones := make([]chan bool, 4)
	//done := make(chan bool)
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ... ?", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done
	}
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
