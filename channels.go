package main

import (
	"fmt"
)

type HttpResponse2 struct {
	ers [][]error
}

func TestChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	b, ok := <-ch
	fmt.Println(b, ok)
	c, ok := <-ch
	fmt.Println(c, ok)
}
