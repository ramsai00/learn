package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go hello(1)
	go hello(2)
	go hello(3)
	go hello(4)
	go hello(5)
	fmt.Println("end")

	time.Sleep(100 * time.Millisecond)
}

func hello(x int) {
	for i := 0; i < 3; i++ {
		fmt.Println(x, ": hello!")
	}
}
