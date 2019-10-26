package main

import "fmt"

func arrayForm() {
	var x = [5]int{12, 13, 14, 15, 16}

	var total float64
	for i := 0; i < 5; i++ {
		total += float64(x[i])
	}

	fmt.Println("----------------------")
	fmt.Println("[SIMPLE ARRAY]")
	fmt.Println("----------------------")
	fmt.Println("Array: ", x)
	fmt.Println("Total: ", total)
}

func arraySpecialForm() {
	var x [5]int
	x[0] = 12
	x[1] = 13
	x[2] = 14
	x[3] = 15
	x[4] = 16

	var total float64
	for _, value := range x {
		total += float64(value)
	}

	fmt.Println("----------------------")
	fmt.Println("[SPECIAL FORM ARRAY]")
	fmt.Println("----------------------")
	fmt.Println("Array: ", x)
	fmt.Println("Total: ", total)
}

func slices() {

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[2:4]
	fmt.Println("----------------------")
	fmt.Println("[SLICES]")
	fmt.Println("----------------------")
	fmt.Println("SLICE: ", x)

	slice1 := make([]float64, 5, 10)
	fmt.Println("SLICE CAPACITY: ", slice1)
}

func appendTest() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 10)

	fmt.Println("----------------------")
	fmt.Println("[APPEND]")
	fmt.Println("----------------------")
	fmt.Println("SLICE 1:", slice1)
	fmt.Println("SLICE 2:", slice2)
}

func copyTest() {
	slice1 := []int{1, 2, 3}
	//slice2 := []int{4, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println("----------------------")
	fmt.Println("[COPY]")
	fmt.Println("----------------------")
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func mapTest() {
	fmt.Println("----------------------")
	fmt.Println("[MAPS / DICTIONARY]")
	fmt.Println("----------------------")
	x := make(map[string]string)
	x["name"] = "zack"
	x["age"] = "35"
	x["country"] = "philippines"

	fmt.Println("Dictionary: ", x)
	delete(x, "age")
	fmt.Println("Dictionary (Deleted Age): ", x)
	fmt.Println("Name: ", x["name"])
	fmt.Println("Age: ", x["age"])

	name, ok := x["name"]
	fmt.Println("Lookup: ", name, ok)

	if name, ok := x["country"]; ok {
		fmt.Println("Lookup: ", name, ok)
	}
}

func complexMaps() {
	fmt.Println("----------------------")
	fmt.Println("[COMPLEX MAPS]")
	fmt.Println("----------------------")
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	if el, ok := elements["H"]; ok {
		println("Elements: ", el["name"], el["state"])
	}

}
