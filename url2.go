package main 
import(
	"fmt"
	"net/http"
	"os"
	"io"
)
func main(){
	for _,url :=range os.Args[1:]{
		var outer io.Writer
		outer,_ = os.OpenFile("./a.txt",os.O_RDWR,0666)
		resp,err := http.Get(url)
		if err !=nil{
			fmt.Println("error")
		}
		data,err :=io.Copy(outer,resp.Body)
		resp.Body.Close()
		fmt.Println(data)
	}
}