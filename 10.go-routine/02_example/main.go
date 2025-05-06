package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	msg := "hello"
	go func(m string) {
		fmt.Println(m)
	}(msg)
	msg = "world"
	fmt.Println("end")
	time.Sleep(time.Millisecond * 100)
}
