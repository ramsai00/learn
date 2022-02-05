package main

import "fmt"

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
	sides         int
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) getSides() int {
	return r.sides
}

func Calculate(s shape) {
	fmt.Println(s.area())
	rect := s.(rectangle)
	fmt.Println(rect.getSides())
}

func main() {
	r := rectangle{
		width:  3,
		height: 4,
		sides:  5,
	}
	Calculate(r)
}
