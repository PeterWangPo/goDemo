package main

import (
	"fmt"
	"reflect"
)

func main() {
	//函数不定参数问题
	test(1, 2, 3)
	fmt.Println("============")
	test(3, 4, 5, 7, 8)
}

//函数不定参数用...标识
func test(args ...int) {
	for _, v := range args {
		fmt.Println(v, reflect.TypeOf(v))
	}
}
