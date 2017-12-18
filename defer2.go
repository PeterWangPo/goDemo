package main

import (
	"fmt"
)

func main() {
	//defer返回一个立刻执行的闭包函数，函数里面的i引用一直存在，最后i的值是3，所以输出都是3
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	//这里会输出210
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
