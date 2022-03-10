package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channels")
	ch := make(chan int, 2) // Buffered channel now it will take size number of parameters excess the size will lead to deadlock
	go receiver(ch)
	ch <- 1
	ch <- 2
	time.Sleep(100 * time.Millisecond)
}

func receiver(ch <-chan int) {
	fmt.Println(<-ch)
}
