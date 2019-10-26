package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readFile() {
	file, err := os.Open("assets/test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println("****** Long way reading a file...")
	fmt.Println(str)
}

func readFileShortWay() {
	bs, err := ioutil.ReadFile("assets/test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println("")
	fmt.Println("****** Short way reading a file...")
	fmt.Println(str)
}

func createFile() {
	file, err := os.Create("assets/created.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("testing creating a file")
	fmt.Println("")
	fmt.Println("****** Writing a file...")
	fmt.Println("assets/created.txt")
}

func getDirectoryContents() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	fmt.Println("")
	fmt.Println("****** Directory contents...")

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func readDirectory() {
	fmt.Println("")
	fmt.Println("****** Read Directory...")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
