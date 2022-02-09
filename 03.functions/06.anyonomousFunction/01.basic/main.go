package main

import "fmt"

/*
func (parameter_list) (return type){
	---code
}()
*/

func main() {
	func() {
		fmt.Println("Welcome")
	}()
	func(ele string) {
		fmt.Println(ele)
	}("Ram")
}
