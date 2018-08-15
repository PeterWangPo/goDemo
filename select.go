package main

import (
	"time"
	"fmt"
)

func main()  {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server2(ch2)
	go server1(ch1)
	select {
	case out1 := <- ch1:
		fmt.Println("output from server:",out1)
	case out2 := <- ch2:
		fmt.Println("output from server:",out2)
	//default:
	//	fmt.Println("none...")
	}
}
func server1(ch chan <- string)  {
	time.Sleep(time.Second*2)
	ch <- "server1"
}
func server2(ch chan <- string)  {
	time.Sleep(time.Second*2)
	ch <- "server2"
}