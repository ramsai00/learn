package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type agent struct {
	person
	licence bool
}

func main() {
	p1 := person{
		first: "Ram",
		last:  "sai",
		age:   21,
	}
	p2 := agent{
		person: person{
			first: "Satya",
			last:  "Venkat",
			age:   21,
		},
		licence: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.person.first, p2.last, p2.age)
}

/*specification as p2.person.first or p2.first based on the embedding useful when other filed
struct consists of same filed name.
*/

//The above agent in person embedded can't be a pointer type
//The above agent embedded person the field of person are called promoted type.
