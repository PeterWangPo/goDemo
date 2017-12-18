package main

import (
	"fmt"
)

type human struct {
	age  int
	name string
}
type student struct {
	human
	ability []string
}
type z struct {
	num   int
	hobby struct {
		add, phone string
	}
}

func main() {
	//嵌入结构声明方式一：key-value形式声明。嵌入的默认key为嵌入结构名称
	p0 := student{human: human{2, "p0"}, ability: []string{"fighting..."}}
	fmt.Println(p0)
	//嵌入结构声明方式二：value形式声明
	p1 := student{human{1, "p1"}, []string{"talk..."}}
	fmt.Println(p1)
	A(&p1)
	fmt.Println(p1)
	fmt.Println("=====================")
	//定义指针类型的struct
	p2 := &student{
		human{3, "p2"},
		[]string{"walking..."}, //如果是这样定义的话，每一行必须要有一个逗号
	}
	fmt.Println(p2)
	A(p2)
	fmt.Println(p2)
	fmt.Println("=====================")
	z1 := z{num: 1} //结构里面套结构定义，那么只能这样声明，然后再下面赋值
	fmt.Println(z1)
	z1.hobby.add = "add1"
	z1.hobby.phone = "phone1"
	fmt.Println(z1)
}

//通过指针传递struct
func A(p *student) {
	p.age = 2
	fmt.Println(p)
}
