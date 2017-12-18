package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("in main ", arr)
}
func modify(arr [5]int) {
	arr[0] = 12
	fmt.Println("in modify", arr)
}
