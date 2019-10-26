package main

import (
	"fmt"
	"math"
)

// Circle struct
type Circle struct {
	x, y, r float64
}

// METHODS
// Since Go always copies arguments,
// If we want to modify on the of fields inside the circleArea function,
// it will not modify the original variable.
// To modify the original variable =, we can use pointer to that circle
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// We no longer need the & operator -
// Go automatically knows to pass a pointer to the circle for this method.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func structTest() {
	c1 := new(Circle)
	var c2 Circle
	c3 := Circle{0, 0, 5}

	fmt.Println(c1) // &{0,0,0}
	fmt.Println(c2) // {0,0,0}
	fmt.Println("")
	fmt.Println("Circle Area 0: ", circleArea(c1)) // doesnt need & since we initialized Circle using "new"
	fmt.Println("Circle Area 1: ", circleArea(&c3))
	fmt.Println("Circle Area 2: ", c1.area())
}
