// A goroutine is an independent function that executes simultaneously in some separate
// lightweight threads managed by Go. GoLang provides it to support concurrency in Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	time.Sleep(1 * time.Second)
	goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("goodbye")
}

// step 1 : main goroutine starts
// step 2 : Invokes helloworld and helloworld goroutine starts
// step 3 : If there is no pause using the sleep method, the main will then invoke goodbye()
// and exit before the helloworld goroutine finishes its execution.
