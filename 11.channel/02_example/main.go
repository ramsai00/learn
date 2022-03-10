package main

import "fmt"

func main() {
	fmt.Println("channels")
	ch := make(chan int)
	go reciever(ch)
	ch <- 40
	fmt.Println(<-ch)
	close(ch)
}

// channels are bidirectional now you can assign values or read values from the channel
func reciever(ch chan int) {
	fmt.Println(<-ch)
	ch <- 1234
	fmt.Println("Channel closed by the sender")
}
