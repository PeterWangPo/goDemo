package main

import (
	"fmt"
)

type human struct {
	age  int
	name string
}
type student struct {
	human   //student继承了human的全部
	ability []string
	age     int
}

func main() {
	li := student{human{11, "lili"}, []string{"study"}, 22}
	fmt.Println(li)
	fmt.Println(li.age)       //如果student本身没有age字段，则访问父级human里面的age，有就访问自己本身的age字段
	fmt.Println(li.human.age) //访问继承的age,可以通过父struct名称的形式访问
	fmt.Println(li.name)
}
