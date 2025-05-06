package main

import "fmt"

func main() {
	// input
	num := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(num)
	// stage 2
	finalChannel := toSquare(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func toSquare(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func sliceToChannel(num []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()

	return out
}
