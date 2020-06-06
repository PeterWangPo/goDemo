package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("in main after modify", arr)
	modefyPoint(&arr)
	fmt.Println("in main after point", arr)
	modifySlice(arr[:])
	fmt.Println("in main slice", arr)
}
func modify(arr [5]int) {
	arr[0] = 12
	fmt.Println("in modify", arr)
}
func modefyPoint(arr *[5]int) {
	arr[0] = 13
	fmt.Println("in modify point", arr)
}

func modifySlice(arr []int) {
	arr[0] = 100
	fmt.Println("in modify slice", arr)
}
