package main 
import "fmt"
func main(){
	var (
	    NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
	)
	fmt.Println(NotPV)
	for _,v := range NotPV{
	    fmt.Println(v)
	}
}
