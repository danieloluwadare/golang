package main

import "fmt"

type Emp struct {
	firstname string
	lastname  string
}

// This method acts like a method of the class/struct Emp
func (e Emp) changeName(newName string) {
	e.firstname = newName
}
func (e *Emp) changeNamePointer(newName string) {
	(*e).firstname = newName
}
func testMethodsOfStructs() {
	emp := &Emp{"Daniel", "Daa"}
	fmt.Println("before name change")
	fmt.Println(emp.firstname)
	emp.changeName("David")
	fmt.Println("after name change")
	fmt.Println(emp.firstname)

	fmt.Println("before name change using changeNamePointer")

	fmt.Println("after name change using changeNamePointer")
	emp.changeNamePointer("David")
	fmt.Println(emp.firstname)
	fmt.Println(emp)

}
