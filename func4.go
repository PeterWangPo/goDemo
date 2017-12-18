package main
import "fmt"
func add(x, y int) (a, b int){
	a = x + y;
	b = x - y;
	//没有参数的return,实际返回的是a,b变量的值。不建议这样写
	return;
}
func main(){
	fmt.Println(add(3, 2));
}