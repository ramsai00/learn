//channel is  a communication object using which go routines can communicate with each other.
// go routines can write data into this pipe or into this go routines and other go routines can read data from another side
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channels")

	ch := make(chan int) // creating a channel
	go receiver(ch)
	ch <- 42 // writing a value into channel
	ch <- 50
	close(ch) // closing a channel
	time.Sleep(100 * time.Millisecond)
}

func receiver(ch chan int) {
	for {
		i, ok := <-ch // Reading a channel
		if ok {
			fmt.Println(i)
		}
	}
}
