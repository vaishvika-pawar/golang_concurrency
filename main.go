package main

import (
	"fmt"
	"time"
)

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)
	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ... ?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		//fmt.Println(doneChan)
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
	close(doneChan)
	// only add this line to the function 
	// which you know will take longest 
	// and after which we don not need the channel
	// <- this operator sends data to channel
}

// Goroutine behavior
// The idea behind running goroutine is to run it in a non-blocking way

// Channels
// A value that can be used as a communication when working with goroutines
