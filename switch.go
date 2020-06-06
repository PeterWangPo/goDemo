package main

import "fmt"

func main() {
	a := 1
	//用法一, case里面不需要写break.会自动跳出
	switch a { //左括号必须和switch在同一行
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("nil")
	}
	fmt.Println("--------------------")

	//用法二：case里面是一个条件判断
	switch {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough //接着执行下面的case，不直接跳出
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("nil")
	}

	//用法三：初始化变量，变量也是块级作用域
	switch b := 1; { //初始化变量后面必须加分号结束
	case b == 0:
		fmt.Println("b == 0")
	case b == 1:
		fmt.Println("b == 1")
	default:
		fmt.Println("nil")
	}
}
