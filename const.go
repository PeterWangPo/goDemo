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

func main() {
	fmt.Println(a, b, c, d, e, f)
}
