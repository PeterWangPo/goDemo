package main 
import "fmt"
func main(){
		count := 0
		for b := 1; b < 3; b++{
			for{
				if count > 3{
					break
				}
				count++
				fmt.Println("label1 here")
				continue
			}
		}
}