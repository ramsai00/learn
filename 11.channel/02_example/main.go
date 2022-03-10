package main

import "fmt"

func main() {
	fmt.Println("channels")
	ch := make(chan int)
	go reciever(ch)
	fmt.Println(<-ch)
	close(ch)
}

func reciever(ch chan int) {
	fmt.Println(<-ch)
	ch <- 1234
	fmt.Println("Channel closed by the sender")
}
