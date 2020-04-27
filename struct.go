package main

import "fmt"

type user struct {
	name    string
	address string
	age     int
}

func testArraysExamplev() {
	fmt.Println("From testArraysExample method")
	var studentsName [3]string
	studentsName[0] = "Daniel"
	studentsName[1] = "Bkm"
	studentsName[2] = "Tomkid"

	fmt.Println(studentsName[1])   //prints Two
	fmt.Println(len(studentsName)) //prints 3
	fmt.Println(studentsName)

	// creating an integer array and the size of the array is defined by the number of elements
	directions := [...]int{1, 2, 3, 4, 5}
	fmt.Print("total number of directions ===>")
	fmt.Println(len(directions)) //prints 3

}

func testReturnArrayExamplec() [3]string {
	var studentsName [3]string
	studentsName[0] = "Daniel"
	studentsName[1] = "Bkm"
	studentsName[2] = "Tomkid"

	return studentsName
}
