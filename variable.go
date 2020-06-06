package main

import "fmt"

var d, q = true, false //有赋值，可不用带变量类型,必须带var
var x, y bool          //只初始化，必须带变量类型，必须带var
func main() {
	var t bool      //只初始化，必须带变量类型，必须带var
	var a, b = 1, 2 //有赋值，可不带变量类型，必须带var
	k, j := 3, 4    //这是函数内的一种特殊变量赋值，简称短赋值，可不用var
	fmt.Println(t, d, q, x, y, a, b, k, j)
}
