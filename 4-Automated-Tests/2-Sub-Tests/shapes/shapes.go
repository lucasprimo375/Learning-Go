package shapes

import (
	"math"
)

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // can be math.Pow(c.radius, 2)
}

type Shape interface {
	Area() float64
}
