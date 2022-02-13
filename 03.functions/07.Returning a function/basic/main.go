package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	s2 := bar()
	fmt.Printf("Type %T, value %v\n", s2, s2)
	s3 := s2()
	fmt.Println(s3)
}

func foo() string {
	s := "Hello World!"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
