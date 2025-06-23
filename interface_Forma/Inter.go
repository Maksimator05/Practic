package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rect := Rectangle{5, 4}
	circle := Circle{5}
	
	fmt.Printf("The area of the circle: %.2f\n", circle.Area())
	fmt.Printf("The area of the rectangle: %.2f\n", rect.Area())
}
