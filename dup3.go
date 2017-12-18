package main 
import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)
func main(){
	files :=os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0{
		fmt.Println("params error")
	}else{
		for _, v := range files{
			data, err := ioutil.ReadFile(v)
			if err != nil {
				fmt.Println("readfile error %v", err)
				continue
			}
			for _,vv :=range strings.Split(string(data),"\r\n"){
				// fmt.Println(vv)
				counts[vv]++
				if counts[vv] >= 2{
					fmt.Println(v)
				}
			}
		}
	}
	// fmt.Println(counts)
}