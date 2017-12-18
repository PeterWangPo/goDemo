package main 
import(
	"net/http"
	"sync"
	"fmt"
)
var mu sync.Mutex
var count int
func main(){
	http.HandleFunc("/", hander)
	http.HandleFunc("/help", helpHander)
	http.ListenAndServe("localhost:8181", nil)
}
func hander(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path: %s", r.URL.Path)
}
func helpHander(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d Method: %s Protol: %s Url: %s", count, r.Method, r.Proto, r.URL)
	mu.Unlock()
}