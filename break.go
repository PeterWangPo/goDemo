package main 
import "fmt"
func main(){
	LABEL1:
		for{
			for {
				fmt.Println(" label1 here")
				break LABEL1    //跳出某个循环,这里跳出到和LABLE1同一层的for循环,也就是最外层的for循环.
			}	
		}
	fmt.Println(" label2 here")
}