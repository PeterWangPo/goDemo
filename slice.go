package main

import (
	"fmt"
)

func main() {
	//array  制定了长度或者使用...自动计算数组长度的，为数组。数组为值传递
	a := [...]int{1, 2}
	b := [2]int{1}
	fmt.Println(a)
	fmt.Println(b)
	//slice 没有指定长度的为slice,其为地址引用,一个变化，其他都会变化
	aSlice := []int{1, 2, 3, 4}
	var bSlice []int //声明一个变量slice，
	aSlice[2] = 6
	aSlice[3] = 5
	//bSlice[1] = 2 //这里居然不能赋值，赋值就报错
	fmt.Println(aSlice)
	fmt.Println(bSlice)
}
