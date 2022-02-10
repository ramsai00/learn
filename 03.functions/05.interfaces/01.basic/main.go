package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licence bool
}

func (s secretAgent) speak() {
	fmt.Println("Hello", s.person.last, s.first, "secret Agent")
}

func (p person) speak() {
	fmt.Println("Hello", p.first, p.last, p.age, "person")
}

type human interface {
	speak()
}

func foo(h human) {
	switch h.(type) {
	case person:
		fmt.Println("person", h.(person).first)
	case secretAgent:
		fmt.Println("secret", h.(secretAgent).first)
	default:
		fmt.Println("default", h)
	}
}

func main() {
	p := secretAgent{
		person: person{
			first: "ram",
			last:  "sai",
			age:   21,
		},
		licence: true,
	}

	p1 := secretAgent{
		person: person{
			first: "satya",
			last:  "venkat",
			age:   25,
		},
		licence: false,
	}

	p2 := person{
		first: "harish",
		last:  "t",
		age:   22,
	}

	p.speak()
	p1.speak()
	p2.speak()

	foo(p)
	foo(p1)
	foo(p2)
}
