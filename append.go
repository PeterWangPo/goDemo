package main

import "fmt"

func sum(num ...int) int {
	ret := 0
	for _, v := range num {
		ret += v
	}
	return ret
}
func main() {
	s1 := []int{1, 2}
	s1 = append(s1, 3)
	fmt.Println(s1)
	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	s1 = append(s1, []int{6, 7, 8}...)
	fmt.Println(s1)
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
