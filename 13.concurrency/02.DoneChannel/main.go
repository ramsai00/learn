package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(10 * time.Second)
}

// done channel we are only passing as read only channel.
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}
