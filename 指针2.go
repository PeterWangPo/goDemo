package main

import "fmt"

func main() {
	a, b := 1, 2
	p := &a         //& 符号会生成一个指向其作用对象的指针
	fmt.Println(*p) //* 符号表示指针指向的底层的值。
	fmt.Println(p)  //
	fmt.Println("=============")
	*p = 3
	fmt.Println(a)
	p = &b
	*p = *p //无用，干扰项
	fmt.Println(b)
	fmt.Println(*p)
}
