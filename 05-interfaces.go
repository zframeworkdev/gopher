package main

import (
	"fmt"
	"math"
)

// Rectangle struct
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Shape interface
type Shape interface {
	area() float64
}

// Variadic arguments
func totalArea(circles ...Circle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	return total
}

// Using the interface in arguments
func totalArea2(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.area()
	}
	return total
}

// MultiShape - INTERFACE AS FIELDS
type MultiShape struct {
	shapes []Shape
}

func interfaceTest() {
	var circle = Circle{1, 2, 5}
	var rectangle = Rectangle{0, 0, 10, 10}
	var rectangleB = new(Rectangle)
	rectangleB.x1 = 1
	rectangleB.y1 = 1
	rectangleB.x2 = 5
	rectangleB.y2 = 5
	total := totalArea2(&circle, &rectangle, rectangleB)
	fmt.Println("Shape Total Area:", total)

	// ....
	interfaceFieldTest()
}

func interfaceFieldTest() {
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
			&Circle{3, 3, 12},
		},
	}
	fmt.Println("Field Shape Total Area:", multiShape.area())
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
