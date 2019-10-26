/*
Avoid rewriting everything from scratch
*/
package main

import (
	"fmt"
	"strings"
)

func testStringsPackage() {
	fmt.Println("Index: ", strings.Index("tested", "e"))

	var str = []string{"a", "b", "c"}
	fmt.Println("Join: ", strings.Join(str, "-"))

	fmt.Println("Repeat: ", strings.Repeat("a", 5))

	fmt.Println("Replace: ", strings.Replace("aaaa", "a", "b", 2))
	fmt.Println("Replace: ", strings.Replace("aaaa", "a", "b", -1))

	var strB = "a-b-c-d"
	fmt.Println("Split: ", strings.Split(strB, "-"))

	stringBytes()
}

func stringBytes() {
	arr := []byte("Spring")
	str := string(arr)

	fmt.Println("String -> Bytes: ", arr)
	fmt.Println("Bytes -> String: ", str)
}
