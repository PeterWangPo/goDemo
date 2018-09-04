package main
import(
	"fmt"
	"os"
)
func main(){
	args := os.Args
	fmt.Println(args)
	for k,v := range args{
		if k == 0{
			continue
		}
		fmt.Println(v)
	}
}