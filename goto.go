package main

import "fmt"

func main() {
	count := 0
LABEL1: //这个词貌似只要不是关键词都可以。。。
	a := 1
	for {
		if a > 3 {
			goto LABLE2 //跳转到LABLE2执行
		}
		fmt.Println(a)
		a++
	}
LABLE2:
	fmt.Println("lable2 here")
	count++
	if count > 3 {
		goto LABEL3
	}
	goto LABEL1
LABEL3:
	fmt.Println("label3 here")
}
