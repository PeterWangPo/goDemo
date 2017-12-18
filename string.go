package main

import (
	"fmt"
)

func main() {
	//字符串直接用+链接
	a, b := "hello", "world"
	fmt.Println(a + b)
	//字符串需要用双引号""或者反引号``包起来，不能使用单引号。``可多行字符串
	c := `
		hello
		world
	`
	fmt.Println(c)
}
