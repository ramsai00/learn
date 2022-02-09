package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	licence bool
}

//func (r receiver) identifier(parameters) (return(s)){code}

func (s secretAgent) speak() {
	fmt.Println("Hello,I am", s.first, s.last)
}

func main() {
	p := secretAgent{
		person: person{
			first: "ram",
			last:  "sai",
		},
		licence: true,
	}
	fmt.Println(p)
	p.speak()
	p1 := secretAgent{
		person: person{
			first: "satya",
			last:  "venkat",
		},
		licence: true,
	}
	p1.speak()
}
