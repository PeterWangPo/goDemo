package main

import (
	"fmt"
)

type Foo struct {
	*Base
	name string
	age  int
}
type Base struct {
	age int
}

func main() {
	foo := Foo{Base: &Base{1}, name: "li", age: 12}
	foo.setAge(3)
	fmt.Println(foo.getAge())
	fmt.Println(foo.Base.getAge())
}
func (this *Base) getAge() int {
	return this.age
}
func (this *Foo) getAge() int {
	return this.age
}
func (this *Foo) setAge(num int) int {
	this.age += num
	return this.age
}
