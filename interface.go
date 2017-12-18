package main
import(
	"fmt"
)
type interger int
func (r interger) less(b interger) bool{
	return r > b
}
func (r *interger) add(b interger){
	*r +=b
}
type a interface{
	less(b interger) bool
	add(b interger)
}
func main(){
	var a1 interger = 1
	var i1 a = &a1
	fmt.Println(i1)
}
