package main
import "fmt"
func add(x int, y int) (int){
	//参数必须给类型，类型在变量名后面。函数返回值也必须定义类型
	return x + y;
}
func main(){
	//main是入口
	fmt.Println(add(1,2))
}