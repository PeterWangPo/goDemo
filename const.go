package main

import (
	"fmt"
)

const (
	a = "A"
	b
	c = iota
	d
	e, f = iota, iota //4,4 iota在增加一行才加一，尽管这又两个，但是只加一，所以两个值都是4
)

const (
	a1 = 'B'
	a2 = iota
	a3
	a4
)

func main() {
	fmt.Println(a, b, c, d, e, f)
	//1.iota只能再const里面使用
	//2.iota再const出现时，就会被重置为0
	fmt.Println(a1, a2, a3, a4)
}
