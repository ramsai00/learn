package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(x...)
	fmt.Println(s)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("pos", i, "add", v, "total", i+v)
	}
	return sum
}
