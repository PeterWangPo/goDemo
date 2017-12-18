package main
import "fmt"
func main(){
  arr:=[6]int{1,2,3,4,5,6}
  s1 :=arr[3:4]
  fmt.Println(s1)
  fmt.Printf("%p",s1)
  s1 = append(s1,1,2,3)
  fmt.Println(s1)
  fmt.Printf("%p",s1)
}
