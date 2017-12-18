package main
import(
	"fmt"
	"os"
)
func main(){
	args := os.Args
	for k,v := range args{
		if k == 0{
			continue
		}
		fmt.Println(v)
	}
	fmt.Println(args[1:2]);
}