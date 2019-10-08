package main

import (
	"fmt"
	"math"
)

// define an interface
type Shape interface {
	area() float64
}

// define a circle struct
type Circle struct {
	x, y, radius float64
}

// define a rect struct
type Rectangle struct {
	height, width float64
}

// func area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// func area of a rectangle
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main()  {
	circle := Circle{0, 0, 7}
	rectangle := Rectangle{4, 5}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}