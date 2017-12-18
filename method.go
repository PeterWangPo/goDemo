package main

import (
	"fmt"
)

type A struct {
	name string
	age  int
}
type Interger int

func main() {
	a := A{"lili", 12}
	fmt.Println(a.getName())
	fmt.Println(a.incAge())
	fmt.Println(a.getAge())
	var i Interger
	i.inc()
	fmt.Println(i)
}
func (this A) getName() string {
	return this.name
}
func (this *A) incAge() int {
	this.age += 3
	return this.age
}
func (this A) getAge() int {
	return this.age
}
func (this *Interger) inc() {
	*this += 3
}
