package main

import "fmt"

func add(a int, b int) int {
	c := a + b
	return c
}

func calculate(a int, b int, f func(int, int) int) int {

	r := f(a, b)
	return r
}

type Calcfunc func(int, int) int

func calculate2(a int, b int, f Calcfunc) int {
	r := f(a, b)
	return r
}

func testFuctionWithFunctionArgument(x int, y int) {
	fmt.Printf("The result = %d  \n", calculate2(x, y, add))

}
