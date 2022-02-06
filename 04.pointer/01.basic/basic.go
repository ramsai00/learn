package main

import "fmt"

// pointer stores the address of the variable

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println("address of a :", &a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Printf("%v %T\n", *b, b)

	*b = 43
	fmt.Println(a)
}

// * gives you the value stored at an address when you have the address
