package main

import "fmt"

// A composite data type is any data type can be constructed using primitive data type

// Aggregate data type is refer to arrays,lists.

func main() {
	//  x := type{values}			composite literals
	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[2]) //access by index.

	for k, v := range x {
		fmt.Println("key :", k, "value:", v)
	}
}
