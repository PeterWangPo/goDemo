package main
import "fmt"
func add(x, y, z int) (int) {
	//多个参数，当参数类型一致时，可以只定义最后一个参数的类型
	return (x + y - z);
}
func main(){
	fmt.Println(add(3, 2, 1));
}