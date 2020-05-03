package main

import "fmt"

func testPointers() {
	var a int = 10
	var pa *int
	pa = &a

	fmt.Printf("pointer pa of type %T with value of %v\n", pa, pa)

	fmt.Println("Derefrencing a pointer => meaning getting the actual value")

	fmt.Printf("pointer pa of type %T with value of %v\n", pa, *pa)

	fmt.Println("Changing the variable value using a pointer")
	*pa = 2

	fmt.Printf("a = %v\n", a)
	fmt.Printf("pointer pa of type %T with value of %v\n", pa, *pa)

	fmt.Println("The new keyword creates a allocates memeory and returns a pointer")
	pb := new(int)
	fmt.Printf("data at memory location %v with value of %v\n", pb, *pb)

	fmt.Println("passing a pointer to a function")
	c := 7

	evaluatePointer(&c)
	fmt.Println("After passing a pointer to a function and changing the pointer value")
	fmt.Printf("data at memory location %v with value of %v\n", c, c)

	fmt.Println("increament of pointers")
	*pb++
	*pb++
	*pb++
	*pb++
	fmt.Printf("increamented pointer  at memory location %v with value of %v\n", pb, *pb)

	// fmt.Println()
	// fmt.Println(compareChar()

}

func evaluatePointer(p *int) {
	*p = 3
}
