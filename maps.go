package main

import "fmt"

func testMapExample() {
	// var m map[string]int // nil map of string-int pairs

	m1 := make(map[string]int) // Empty map of string-int pairs
	// m2 := make(map[string]int, 100) // Preallocate room for 100 entries

	m3 := map[string]float64{ // Map literal
		"e":  2.71828,
		"pi": 3.1416,
	}
	fmt.Println(len(m3)) // Size of map: 2

	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3

	fmt.Printf("The length of map m1 is  = %d \n", len(m1))

	// get value of entry three;
	fmt.Printf("The key one of map m1 is  = %d \n", m1["one"])

	_, found := m1["pi"]
	fmt.Printf("The found variable  = %t \n", found)

	if x, found := m1["three"]; found {
		fmt.Println(x)

	}

	for key, value := range m1 { // Order not specified
		fmt.Println(key, value)
	}

}

// When you return a slice you dont need to specify the length of the slice
func testMapExampleReturn() map[string]float64 {
	m3 := map[string]float64{ // Map literal
		"e":  2.71828,
		"pi": 3.1416,
	}
	return m3
}
