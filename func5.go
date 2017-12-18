package main

import (
	"fmt"
)

func main() {
	//函数引用传值
	//默认情况下，一般数据类型函数参数都是形参，传递的是值
	//而channel,slice,map传递的是引用
	a := 1
	fmt.Println(add(&a)) //调用函数的时候需要引用传递
	p := &a              //引用赋值
	fmt.Println(p)       //只输出p的地址
	fmt.Println(*p)      //输出p的值
}
func add(x *int) int { //申明函数的时候参数需要引用传递
	t := *x + 1
	return t
}
