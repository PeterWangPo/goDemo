package main

import "fmt"

func main() {
	//只初始化，未赋值，就会被设置为零值。零值包括0,false,''
	var i int
	var m bool
	var d string
	var c float64
	fmt.Printf("%v, %v, %v, %v", i, m, d, c)
}
