package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Calculate(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{
		width:  5,
		height: 6,
	}
	c := circle{
		radius: 6,
	}
	Calculate(r)
	Calculate(c)
}
