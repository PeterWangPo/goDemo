package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	//声明一：按照key-value声明,不用保证顺序
	a := person{
		name: "lili",
		age:  20,
	}
	//struct是值传递，只有map,slice,channel是引用传递
	fmt.Println("a:", a)
	A(a)
	fmt.Println("a:", a)
	//声明二：按照value声明，不过要保证顺序
	b := person{"meimei", 11}
	fmt.Println(b)
}
func A(per person) {
	per.age = 18
	fmt.Println("per :", per)
}
