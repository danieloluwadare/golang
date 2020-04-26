package main

import "fmt"

func testSliceExample() {
	studentSlice := make([]string, 0) // same as []int{0, 0}
	studentSlice = append(studentSlice, "Hello")
	studentSlice = append(studentSlice, "World")

	for i, s := range studentSlice {
		fmt.Println(i, s)
	}
	// creating an integer slice
	var directions []int
	directions = append(directions, 6)

	fmt.Print("total number of directions ===>")
	fmt.Println(len(directions)) //prints 1

	retSlice := testSliceExampleReturn()
	fmt.Print("Returned Slice ===>")
	fmt.Println(retSlice) //prints

}

// When you return a slice you dont need to specify the length of the slice
func testSliceExampleReturn() []string {
	studentSlice := make([]string, 0) // same as []int{0, 0}
	studentSlice = append(studentSlice, "Hello")
	studentSlice = append(studentSlice, "World")

	for i, s := range studentSlice {
		fmt.Println(i, s)
	}
	return studentSlice
}
