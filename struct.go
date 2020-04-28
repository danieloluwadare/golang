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

}
