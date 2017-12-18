package main

import (
	"fmt"
)

type skill []string
type human struct {
	age  int
	name string
}
type student struct {
	human
	skill
	int
	special string
}

func main() {
	s1 := student{human: human{1, "li"}, special: "study"}
	s1.int = 1
	fmt.Println(s1.int)
	s1.skill = []string{"abc", "dev"}
	fmt.Println(s1.skill)
}
