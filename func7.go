package main

import (
	"fmt"
)

func main() {
	j := 10
	//闭包函数 匿名函数闭包
	a := func() func() {
		i := 5
		return func() {
			fmt.Println(i, j)
		}
	}()
	a()
	j = j * 2
	a()
	fmt.Println("===============")
	m := closure(3)
	fmt.Println(m(1))
	fmt.Println(m(2))
}

func closure(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
