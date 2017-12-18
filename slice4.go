package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	s1 = append(s1, 5)
	s2 := []int{7, 8}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	s1 = s1[:3]
	fmt.Println(s1)
	fmt.Println(cap(s1))
	s3 := make([]int, 4, 8)
	fmt.Println(s3)
	s4 := []int{1, 2, 3, 4, 5}
	s5 := s4[:2]
	fmt.Println(s5)
	s4[1] = 9
	fmt.Println(s5)
}
