// In concurrent programming, Go provides channels that you can use for
// bidirectional communication between goroutines.

// Bidirectional communication means that one goroutine will send a message and the other will read it.
// Sends and receives are blocking. Code execution will be stopped until to write and read are done successfully.

// There are a couple different types of channels:

// Unbuffered channel: Unbuffered channels require both the sender and receiver to be present to be
// successful operations. It requires a goroutine to read the data, otherwise, it will lead to deadlock.
// By default, channels are unbuffered.

// Buffered channel: Buffered channels have the capacity to store values for future processing.
// The sender is not blocked until it becomes full, and it doesn't necessarily need a reader to complete
// the synchronization with every operation.

// If a space in the array is available, the sender can send its value to the channel and complete its send operation immediately.

// After its execution, if a receiver comes, the channel will start sending values to the receiver
// and it will start its operation once it receives the values. As the sender and receiver are operating
// at different times, this is called asynchronous communication.

// Syntax to declare a channel
// ch := make(chan Type)

// Declaration of channels based on directions
// 1. Bidirectional channel : chan T
// 2. Send only channel: chan <- T
// 3. Receive only channel: <- chan T

// How to write to channel.
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go greet(msg)
	time.Sleep(2 * time.Second)

	greeting := <-msg

	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Ramsai"

	fmt.Println("Greeter Channel")
}
