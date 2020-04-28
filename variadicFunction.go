package main

import (
	"fmt"
)

func testVariadicFunction() {
	var sum int = variadicFunction(1, 2, 3, 4, 5)
	fmt.Println(sum)

}

// When you return a slice you dont need to specify the length of the slice
func variadicFunction(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}
