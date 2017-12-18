package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	var b []int
	b = a[2:6]
	for _, v := range b {
		fmt.Println(v)
	}
	c := make([]int, 3, 6)
	s1 := append(c, 1, 2, 3)
	fmt.Println(s1)
	fmt.Printf("%p\n", s1)
	s2 := append(c, 1, 2, 3, 4)
	fmt.Println(s2)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s1)
}
