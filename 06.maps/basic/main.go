package main

import "fmt"

// declaration of a map
var employee = map[string]int{"ram": 21, "sundar": 22, "likitha": 21}

// empty map declaration
var emp = map[string]int{}

// The make function takes as argument the type of the map and it returns an initialized map.
var employees = make(map[string]int)

func main() {
	fmt.Println(employee)
	fmt.Println(emp)
	fmt.Println(employees)
	employees["mark"] = 11
	employees["sun"] = 21
	employees["sundar"] = 22
	fmt.Println(employees)
	fmt.Println(len(employees)) // to know the length of the map.

	// to iterate over a map
	for i, v := range employees {
		fmt.Println("KEY :", i, ",VALUE :", v)
	}
}
