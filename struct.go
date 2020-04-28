package main

import "fmt"

// "User...User"
type User struct {
	name     string
	address  string
	age      int
	posatble bool
}
type Address struct {
	nameOfResidence string
	number          int
}

type UserAddress struct {
	name     string
	address  Address
	age      int
	posatble bool
}

func testStructs() {
	fmt.Println("From testArraysExample method")

	var daniel User
	fmt.Println(daniel)

	// setting a struct property
	var david User
	david.age = 20
	david.name = "david"
	david.address = "Not found"
	david.posatble = true
	fmt.Println(david)

	// getting a struct property
	fmt.Println(david.posatble)

	// Initializing a struct
	// Here oredr of initializing doesnt matter
	// its not also cumpulsory to initialize all the fields
	debby := User{
		name:     "debby",
		address:  "Not",
		age:      7,
		posatble: true,
	}
	fmt.Println(debby)

	// Initializing a struct without using the field name
	josh := User{"Josh", "Mot", 1200, true}
	fmt.Println(josh)

	// Initializing a Nested Struct

	lekan := UserAddress{
		name:     "lekan",
		address:  Address{"Olo", 10},
		age:      7,
		posatble: true,
	}
	fmt.Println(lekan)

	// A nested struct property automatically is inherited by the Parent Struct

	// var studentsName [3]string
	// studentsName[0] = "Daniel"
	// studentsName[1] = "Bkm"
	// studentsName[2] = "Tomkid"

	// fmt.Println(studentsName[1])   //prints Two
	// fmt.Println(len(studentsName)) //prints 3
	// fmt.Println(studentsName)

	// // creating an integer array and the size of the array is defined by the number of elements
	// directions := [...]int{1, 2, 3, 4, 5}
	// fmt.Print("total number of directions ===>")
	// fmt.Println(len(directions)) //prints 3

}
