package main

import "fmt"

func main() {
	myChannel := make(chan string)
	myChannel1 := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		myChannel1 <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)

	case msgFromMyChannel1 := <-myChannel1:
		fmt.Println(msgFromMyChannel1)

	}
}
