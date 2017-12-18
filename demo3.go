package main

import (
	"fmt"
)

var a = b + c //全局变量的声明和顺序无关
var b = 1
var c = 2

func main() {
	fmt.Println(a)
	//一个if elseif声明的变量，在if或者elseif,else里面或者条件里面都是可访问的
	if x := 1; x > 2 {
		fmt.Println("X > 2")
	} else if y := 3; y > 4 {
		fmt.Println("Y > 4")
	} else if y > x {
		fmt.Println("Y > X")
		fmt.Println(y, x)
	}
}
