package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius // can be math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func calculateArea(s shape) {
	fmt.Printf("The area of the shape is %0.2f\n", s.area())
}

func main() {
	fmt.Println("Interfaces")

	r := rectangle{5, 10}
	calculateArea(r)

	c := circle{10}
	calculateArea(c)
}
