package main

import "fmt"

func add(x, y, z int) (int, int) {
	//函数多返回值
	return x + z, y + z
}
func main() {
	fmt.Println(add('a', 'b', 'c'))
	fmt.Println(add(1, 2, 3))
	a, b := add(2, 3, 4)
	fmt.Println(a, b)
}
