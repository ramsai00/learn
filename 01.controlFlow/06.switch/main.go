package main

import "fmt"

func main() {
	x := 3
	switch {
	case x <= 0:
		fmt.Println("false")
	case x == 4:
		fmt.Println("not true")
	case x == 3:
		fmt.Println("true")
		fallthrough
	default:
		fmt.Println("default")
	}
}
