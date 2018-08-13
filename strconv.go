package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.IntSize)
	val, err := strconv.ParseBool("t")
	fmt.Println(val, err)
	val2, err2 := strconv.ParseInt("1111", 2, 8)
	fmt.Println(val2, err2)
}
