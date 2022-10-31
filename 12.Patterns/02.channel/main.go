package main

import "fmt"

func main() {
	myChannel := make(chan string) // unbuffered channel

	// buffered channel allows communication sync only.

	go func() {
		myChannel <- "data"
	}()

	result := <-myChannel

	fmt.Println(result)
}
