package main

import "fmt"

func testIfStatements() {
	fmt.Println("From testIfStatements method")
	var sum int = 6

	if sum > 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}

func testSwitchCaseStatemets() {
	fmt.Println("From testSwitchCaseStatemets method")
	var s int = 2

	switch s {
	case 1:
		fmt.Printf("The s = %d \n", 1)

	case 2:
		fmt.Printf("The s = %d \n", 2)

	case 3:
		fmt.Printf("The s = %d \n", 3)

	default:
		fmt.Println("Printing default")

	}

}
