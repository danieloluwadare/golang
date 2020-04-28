package main

import "fmt"

func testString() {
	var item string = "Hello world"
	var item2 string = "Hello people"

	var itemCharh rune = 'h'
	var itemCharJ rune = 'j'

	s := []rune(item)
	for i := 0; i < len(item); i++ {
		fmt.Printf("%c", item[i])
	}

	fmt.Println("")

	for i := 0; i < len(item); i++ {
		fmt.Printf("%v ", item[i])
	}

	fmt.Println("Rune ======> ")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	fmt.Println("")

	fmt.Println(compareChar(itemCharJ, itemCharh))
	fmt.Println(compareChar(itemCharJ, itemCharJ))

	fmt.Println(item == item)
	fmt.Println(item == item2)

}

func compareChar(x rune, y rune) bool {
	if x == y {
		return true
	}
	return false
}
