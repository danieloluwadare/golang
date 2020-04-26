package main

import (
	"fmt"
)

func testDataTypeExample() {
	//declaring a integer variable x
	var variable1 int
	variable1 = 3                        //assigning variable1 the value 3
	fmt.Println("variable1:", variable1) //prints 3

	//declaring a integer variable variable2 with value 20 in a single statement and prints it
	var variable2 int = 20
	fmt.Println("variable2:", variable2)

	//declaring a variable z with value 50 and prints it
	//Here type int is not explicitly mentioned
	var z = 50
	fmt.Println("z:", z)

	//Multiple variables are assigned in single line- i with an integer and j with a string
	var i, j = 100, "hello"
	fmt.Println("i and j:", i, j)

	//Note that You used := instead of =. You cannot use := just to assign a value to a
	//variable which is already declared. := is used to declare and assign value.
	a := 20
	fmt.Println(a)
}
