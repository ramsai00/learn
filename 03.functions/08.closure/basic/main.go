package main

import "fmt"

func main() {
	GFG := 0

	counter := func() int {
		GFG += 1
		return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())

	a := incrementor()
	fmt.Println(a)

	b := a()
	fmt.Println(b)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
