package main

import (
	"fmt"
)

func main() {
	//一个utf8占三个字节
	a := "helllo world 你好"
	l := len(a)
	for i := 0; i < l; i++ {
		t := a[i]
		fmt.Println(i, t)
	}
	fmt.Println("========================")
	arr := [...]int{22, 33, 44, 55, 66, 77, 88, 99}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("====================")
	brr := arr[0:6]
	for _, v := range brr {
		fmt.Println(v)
	}
}
