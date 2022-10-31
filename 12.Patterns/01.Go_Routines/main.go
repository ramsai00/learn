package main

import (
	"fmt"
	"time"
)

func main() {
	go someFunc("1") // async call
	go someFunc("2")
	go someFunc("3")

	time.Sleep(2 * time.Second)

	fmt.Println("Hi")
}

func someFunc(x string) {
	fmt.Println(x)
}
