package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b []int
	b = a[:3]
	fmt.Println(b)
	fmt.Println(28 * 365)
}
