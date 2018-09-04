package main

import "fmt"

func main()  {
	s1 := "abcdefgSeñor"
	for k, v := range s1 {
		fmt.Println(k,v)
		fmt.Println(k,string(v))
	}
	for i:=0 ; i< len(s1); i++{
		fmt.Println(s1[i])
	}
}
