package main

import "fmt"

// this is a comment

func main() {
	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** COLLECTIONS")
	fmt.Println("******************************")
	arrayForm()
	arraySpecialForm()
	slices()
	appendTest()
	copyTest()
	mapTest()
	complexMaps()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** FUNCTIONS")
	fmt.Println("******************************")
	average()
	myNumber := myNumber(30)
	fmt.Println("Return Function: ", myNumber)
	fmt.Println("")

	variadic := variadicFunction(1, 2, 3)
	fmt.Println("Variadic Function: ", variadic)

	array := []int{1, 2, 3}
	variadic2 := variadicFunction(array...)
	fmt.Println("Variadic Function (From Array): ", variadic2)
	fmt.Println("")

	nextEven := makeEvenGenerator()
	fmt.Println("Closure Result: ", nextEven()) // 0
	fmt.Println("Closure Result: ", nextEven()) // 2
	fmt.Println("Closure Result: ", nextEven()) // 4
	fmt.Println("")

	factorialResult := factorial(5)
	fmt.Println("Factorial Result: ", factorialResult)
	fmt.Println("")

	fmt.Println("Defer Result: ")
	deferSample()
	panicRecoverSample()

	fmt.Println("")
	pointersTest()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** STRUCT")
	fmt.Println("******************************")
	structTest()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** METHODS")
	fmt.Println("******************************")
	methodTest()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** INTERFACE")
	fmt.Println("******************************")
	interfaceTest()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** STRINGS")
	fmt.Println("******************************")
	testStringsPackage()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** OS - Input/Output")
	fmt.Println("******************************")
	readFile()
	readFileShortWay()
	createFile()
	getDirectoryContents()
	readDirectory()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** ERRORS")
	fmt.Println("******************************")
	errorLessons()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** LINKED LIST")
	fmt.Println("******************************")
	listLesson()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** SORT")
	fmt.Println("******************************")
	sortLesson()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** HASH AND CRYPTOGRAPHY")
	fmt.Println("******************************")
	hashLesson()
	compareFilesHash()
	sha1Test()

	fmt.Println("")
	fmt.Println("******************************")
	fmt.Println("****** TCP SERVER")
	fmt.Println("******************************")
	tcpTest()

	fmt.Println("")
}
