package main

import (
	"fmt"
)

func main() {
	var a interface{}
	var i = 5
	a = i
	if value, ok := a.(int); ok {
		fmt.Println("int:", value)
	} else {
		fmt.Println("not int")
	}
}
