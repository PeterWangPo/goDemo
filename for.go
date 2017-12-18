package main
import "fmt"
func main(){
	a := 1
	for { //没有任何条件的for
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("-------------------")

	b := 1
	for b < 3 { //有条件的for
		fmt.Println(b)
		b++
	}
	fmt.Println("-------------------")

	for c := 1; c < 3; c++ { //有初始化和条件的for  和if语句一样，左括号必须和for在同一行。c是块级作用域
		fmt.Println(c)
	}
}