package main

import "fmt"

func main() {
	a := 10
	//if语句不能加括号,支持块级作用域，if里面申明了变量a,那边该变量只能在if或else里面使用
	if a := 3; a > 0 { //左括号必须和if在同一行否则就会编译错误
		fmt.Println(a)
	} else { //同理，括号必须在else同一行。
		fmt.Println("-xxxxxx")
	}
	fmt.Println(a)
	fmt.Println("----------------------------")
	b := 3
	if b > 0 {
		fmt.Println(b)
		fmt.Println("b > 0")
	}
	m := test()
	fmt.Println(m)
}

func test() int {
	a := [...]int{1, 2, 3, 4, 5}
	if a[1] == 2 {
		return a[1]
	} else {
		return a[2]
	}
}
