package main

import "fmt"

func main() {
	var x [5]int
	x = [5]int{1, 2, 3, 4, 5}
	fmt.Print(x)
	x[3] = 42
	fmt.Println(len(x))
}
