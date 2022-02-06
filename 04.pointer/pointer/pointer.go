package main

import "fmt"

/* when we have large amount of data it is good to pass
with the pointer.

If we need to change the */

func main() {
	x := 0
	foo(x)
	pointer(&x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func pointer(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}
