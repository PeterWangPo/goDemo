package main

import (
	"fmt"
)

func main() {
	a := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{22, 33, 44, 55}}
	b := [2][4]int{{1, 2, 3, 4}, {22, 33, 44, 55}}
	fmt.Println(a == b)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
