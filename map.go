package main

import (
	"fmt"
)

func main() {
    t1 := []int{1,2,2}
    all := make(map[int]struct{})
    exists := make(map[int]struct{})
    for _, v := range t1 {
        if _, ok := all[v]; ok {
            exists[v] = struct{}{} 
        } else {
            all[v] = struct{}{}
        }
    }
    fmt.Println("all:",all)
    fmt.Println("exists:",exists)
    for k, _ := range all {
        if _, ok := exists[k]; ok {
            delete(all, k)
        }
    }
    fmt.Println(all)
	var a map[int]string //声明方式一，初始化
	fmt.Println(a)
	b := map[int]string{} //声明方式二，直接赋值
	fmt.Println(b)
	c := make(map[int]string, 2)
	fmt.Println(c) //声明方式三，通过make声明

	//赋值
	c[1] = "hello"
	fmt.Println(c)
	//取值
	fmt.Println(c[1])
	//值不存在呢？为空
	fmt.Println("访问不存在的值", c[2])
	//删除map指定key
	delete(c, 1) //第一个参数是要删除的map名称，第二个参数是要删除的key
	fmt.Println("删除指定值", c[1])
	fmt.Println("==========================")
	m := make(map[int]string, 4)
	m[1] = "a1"
	m[3] = "a3"
	fmt.Println("m values:", m)
	//遍历
	for v, ok := range m {
		fmt.Println(v, ok)
	}
	//判断某个key是否存在,如果1这个key存在，ok就返回true
	value, ok := m[1]
	if ok {
		fmt.Println(value)
	}
	if value2, ok2 := m[2]; ok2 {
		fmt.Println(value2)
	} else {
		fmt.Println("not exists")
	}
}
