package main

import (
	"fmt"
	"time"
)

func main() {
	//可以这样理解，函数内有一个栈，当申明一个defer的时候，就入栈，当函数执行完后，在执行defer栈里面的数据，先进的后出执行
	defer fmt.Println("aaaaa") //不会立刻调用，带上层函数执行完后 按照defer栈先进后出的循序调用
	fmt.Println("bbbbbb")
	fmt.Println(time.Now())

	fmt.Println("loop begin....")
	for i := 1; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("loop done")
	fmt.Println("====================")
}
