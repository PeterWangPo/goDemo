package main

import (
	"fmt"
	"math"
	// "strconv"
)

func main() {
	total := 33333
	v := float64(total)
	// if er != nil {
	// 	panic(er)
	// }
	size := int(math.Ceil(v / 4)) //8334
	num := total / size
	fmt.Println("总页数:", num)
	fmt.Println("每页数据量:", size)
	for i := 0; i < num; i++ {
		if i == 0 {
			fmt.Println(0, size)
		} else {
			fmt.Println(size*i+1, size)
		}
	}
}
