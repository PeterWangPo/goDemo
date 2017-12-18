package main

import (
	"fmt"
)

func main() {
	t1("1", t2)
}
func t1(a string, b func()) {
	fmt.Println(a)
	b()
}
func t2() {
	fmt.Println("t2")
}
