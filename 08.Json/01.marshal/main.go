package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Ram",
		Last:  "sai",
		Age:   21,
	}

	p2 := person{
		First: "satya",
		Last:  "venkata",
		Age:   21,
	}

	people := []person{p1, p2}

	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
