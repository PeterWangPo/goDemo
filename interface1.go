package main
import(
	"fmt"
)
type a interface{}
func main(){
	var a1 a = "1"
	switch a1.(type){
		case int :
			fmt.Println("int")
		case string:
			fmt.Println("string")
	}
}
