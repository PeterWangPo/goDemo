package main

import "fmt"

func main() {
	var nums [5][0]int
	fmt.Println(nums)
	for range nums {
		fmt.Println("hell0")
	}
	var a [2]int
	var b [2]int
	fmt.Println(a)
	fmt.Println(a == b)
	fmt.Println("===============")
	c := [2]int{1}      //第一个元素值为1，其他默认给零
	d := [20]int{19: 1} //第20个元素值为1，其他默认为零
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("================")
	e := [...]int{2: 1, 4: 2, 8: 1} //...表示数组长度不确定，根据赋值自动判断
	fmt.Println(e)
	m := 1
	var p *int = &m //指针
	fmt.Println(p)
	fmt.Println(*p)
	n := [...]int{1}
	fmt.Println("=================")
	fmt.Println(n)
	var g *[1]int = &n //数组指针
	fmt.Println(g)
	fmt.Println(*g)
	var j = new([2]string)
	j[1] = "aa"
	fmt.Println(j)
}
