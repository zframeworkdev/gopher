package main

import (
	"errors"
	"fmt"
)

func errorLessons() {
	err := errors.New("This is how to create an error with a message")
	fmt.Println(err)
}
