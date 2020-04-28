package main

import "fmt"

type IShape interface {
	Area() int
	Perimeter() int
}

// "User...User"
type Rect struct {
	width  int
	height int
}

func (r Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Perimeter() int {
	return r.width * r.height
}

func testInterface() {
	var s IShape
	s = Rect{3, 3}
	r := Rect{2, 2}

	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("Area of rectangle s", s.Area())
	fmt.Println("Area of rectangle r", r.Area())

	fmt.Println("s == r is", s == r)

}
