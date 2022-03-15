package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Address string
}

func New(name string, age int, address string) person {
	return person{
		Name:    name,
		Age:     age,
		Address: address,
	}

}

func main() {
	fmt.Println(New("ram", 21, "RJY"))
}
