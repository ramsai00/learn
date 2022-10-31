package main

import "fmt"

func main() {
	charChannel := make(chan string, 3) // buffered channel
	chars := []string{"a", "b", "c"}

	for _, v := range chars {
		select {
		case charChannel <- v:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
