package main

import (
	"errors"
	"fmt"
)

func  myerror() error {
	return errors.New("let us know if error")
}

func TestError()  {
	err:=myerror()
	if err!=nil{
		fmt.Printf("The error result  = %s ", err.Error())

	}else {
		fmt.Printf("The No error result")

	}
}