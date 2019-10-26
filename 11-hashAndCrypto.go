package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// A hash function takes a set of data and reduces it to a smaller fixed size.
// hashes are frequenty used in programming for everuthing from looking up data to easily detecting changes.
// cryptographic and non

// 32-bit hash
func hashLesson() {
	// create a hasher
	h := crc32.NewIEEE()

	// write our data to it
	h.Write([]byte("test"))

	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}

func getHash(filename string) (uint32, error) {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	// remember to always close opened files
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()

	// copy the file in to the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)

	// we don;t care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func compareFilesHash() {
	h1, err := getHash("assets/test.txt")
	if err != nil {
		return
	}

	h2, err := getHash("assets/created.txt")
	if err != nil {
		return
	}
	fmt.Println("Hash File Results.......... ")
	fmt.Println("File 1 assets/test.txt: ", h1)
	fmt.Println("File 2 assets/created.txt: ", h2)
	fmt.Println("Files are the same?: ", h1 == h2)
}

// 160-bit hash
func sha1Test() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})

	fmt.Println("")
	fmt.Println("SHA 1 Results.......... ")
	fmt.Println(bs)
}
