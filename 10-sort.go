package main

import (
	"fmt"
	"sort"
)

// -----------------------------
// SORT
// -----------------------------

// Alien here...
// -----------------------------
type Alien struct {
	name string
	age  int
}

// ByName here...
// -----------------------------
type ByName []Alien

func (ps ByName) Len() int {
	return len(ps)
}

// defines the ordering
func (ps ByName) Less(i, j int) bool {
	return ps[i].name < ps[j].name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func sortLesson() {
	kids := []Alien{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByName(kids))

	fmt.Println("Sort Result: ", kids)
}
