package main

import "fmt"

//03.functions are all about modular it is all about breaking into small types.

/* 01.syntax
func (r receiver) identifier(parameter) (return){
 	-----
	-----
}
*/

func main() {
	foo()
	bar("ram")                           //marshal
	s1 := simple("hi")                   //argument
	fmt.Println(s1)                      //return
	x, y := mulitiple("satya", "venkat") //multiple return
	fmt.Println(x, y)
}

// when we define a functions the value we pass is called parameter the function we are calling the value given is argument

func foo() {
	fmt.Println("Call from foo")
}

func bar(s string) {
	fmt.Println("Hello ", s)
}

//Everything in GO is pass by value Only.

func simple(s string) string {
	return fmt.Sprint("Hello from simple, ", s)
}

func mulitiple(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, "Hello")
	b := false
	return a, b
}
