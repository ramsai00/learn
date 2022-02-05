package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	y := []int{7, 8, 9, 10}
	x = append(x, y...) // append a slice to slice
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // deleting from a slice
	fmt.Println(x)
}
