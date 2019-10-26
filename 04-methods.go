package main

import "fmt"

// METHODS.....

// Person struct
// -----------------------------
type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

// Color struct
// -----------------------------
type Color struct {
	name string
}

func (c *Color) color() {
	fmt.Println("My color is", c.name)
}

// EMBEDDED TYPES.....

// Android struct (Anonymous Field - Color)
// -----------------------------
type Android struct {
	person Person
	model  string
	Color
}

// -----------------------------
func methodTest() {
	a := new(Android)
	a.person.name = "ZALDY"
	a.person.talk()
	// can call method of an embedded type directly since we just use "Color"
	a.color()
}
