package main

import "fmt"

//variadic functions are the functions which can be taken any number of parameters
//variadic return type of slice Datatype.

func main() {
	foo(2, 3, 4, 5, 6, 7)
	result := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are adding", v, "to the total is ", i+v)
	}
	return sum
}
