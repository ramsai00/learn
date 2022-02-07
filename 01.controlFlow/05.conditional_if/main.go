package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("1.", x)
	} else if x < 40 {
		fmt.Println("2.", x)
	} else {
		fmt.Println("3.", x)
	}
}
