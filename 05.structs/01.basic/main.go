package main

import "fmt"

//values of different typed
//Composite dataTypes Aka Aggregate datatype

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Ram",
		last:  "sai",
	}
	p2 := person{
		first: "satya",
		last:  "Venkat",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
