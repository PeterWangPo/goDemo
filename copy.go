package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{7, 8}
	copy(a, b) //将bcopy到a,索引相同的就覆盖
	fmt.Println(a)
	d := []int{1, 2, 3, 4, 5}
	e := []int{7, 8}
	copy(e, d)
	fmt.Println(e)
}
