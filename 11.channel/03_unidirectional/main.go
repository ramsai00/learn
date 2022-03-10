package main

import "fmt"

func main() {
	fmt.Println("channels")
	ch := make(chan int)
	go receiver(ch)
	ch <- 40
	go sender(ch)
	fmt.Println(<-ch)
	close(ch)
}

// now channel will receive values happen in unidirectional but assign values will error
func receiver(ch <-chan int) {
	fmt.Println(<-ch)
	fmt.Println("Channel closed by the sender")
}

// with the change in the place of arrow sign we can now assign values to channel but reading values will error
func sender(ch chan<- int) {
	ch <- 1234
}
