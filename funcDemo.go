package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer1:", i)
		defer func() {
			fmt.Println(i)
		}()
		fs[i] = func() { fmt.Println("fs:", i) }
	}
	for _, f := range fs {
		f()
	}
}
