package main

import "fmt"

func average() {
	//panic("Not Implemented")
}

func myNumber(r int) int {
	return r
}

// Variadic
func variadicFunction(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

// Closure
func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Defer
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func deferSample() {
	defer second()
	first()

	/*
		defer is often used when resources need to be freed in some way. For example, when we open a file, we need to make sure to close it later.
		With defer:

		f, _ := os.Open(filename)
		defer f.Close()

		This has three advantages:
		1. It keeps our Close call near our Open call so it’s easier to understand.
		2. If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
		3. Deferred functions are run even if a runtime panic occurs.
	*/
}

// Panic & Recover (cause a runtime error)
// panic("Not Implemented")
func panicRecoverSample() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

// Pointers
func zero(xPtr *int) {
	*xPtr = 1000
}
func pointersTest() {
	x := 5
	zero(&x)
	fmt.Println("Pointers Test: ", x)
}
