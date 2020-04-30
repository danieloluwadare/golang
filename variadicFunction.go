package main

import (
	"fmt"
)

func testVariadicFunction() {
	var sum int = variadicFunction(1, 2, 3, 4, 5)
	fmt.Println(sum)

}

func variadicFunction(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}
