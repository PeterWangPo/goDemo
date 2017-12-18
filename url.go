package main 
import(
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)
func main(){
	urls :=os.Args[1:]
	for _,url := range urls{
		resp, err := http.Get(url)
		if err !=nil{
			fmt.Println("error %v",err)
		}
		data,err1 :=ioutil.ReadAll(resp.Body)
		if err1 != nil{
			fmt.Println("error %v", err1)
		}
		resp.Body.Close()
		fmt.Printf("%s",data)
	}
}