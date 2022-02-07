package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
	fmt.Println(loop(5))
}

func loop(x int) bool {
	if x < 5 {
		return false
	}
	return true
}
