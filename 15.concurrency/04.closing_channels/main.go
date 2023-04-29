// Closing the channel: Closing the channel indicates that no more values should be sent on it.
// We want to show that the work has been completed and there is no need to keep a channel open.

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

	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Ramsai"

	close(ch)

	fmt.Println("Greeter completed")
}
