package main

import "fmt"

//type person struct {
//	first string
//	last  string
//	age   int
//}
//
//func main() {
//	p1 := person{
//		first: "james",
//		last:  "Bond",
//		age:   21,
//	}
//	fmt.Println(p1)
//}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "ram",
		last:  "sai",
		age:   21,
	}
	fmt.Println(p1)
}

//The above code is an example of an anonymous datatype because it contain name
