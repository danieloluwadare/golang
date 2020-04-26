package main

import "fmt"

func testForLoopsExamples() {
	var sum int = 0

	for i := 1; i < 5; i++ {
		sum += i
	}

	// The init statement, i := 1, runs.
	// The condition, i < 5, is computed If true, the loop body runs, otherwise the loop is done.
	// The post statement, i++, runs.

	fmt.Println(sum)
	// 10 (1+2+3+4)

}

func testWhileLoopsExamples() {
	var sum int = 0
	i := 0
	// If you skip the init and post statements, you get a while loop.

	for i < 5 && sum < 7 {
		sum += i
		i++
		fmt.Printf("The sum = %d \n", sum)
		fmt.Printf("The iterator = %d \n", i)
	}
	fmt.Println(sum) // 10 (1+2+3+4)

}
