package main
import "fmt"
const (
	n5, n4 = "string", 12
)
func main(){
	var (
		n1 , n2 = len(n5), (n4 + 1);
	);
	const n3 = "a"
	fmt.Println(n1, n2, n3, n4, n5);
}