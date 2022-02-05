package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println("1: ", x)
	fmt.Println("2: ", x[1])
	fmt.Println("3: ", x[1:])
	fmt.Println("4", x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
