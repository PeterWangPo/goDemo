package main

import (
	"sync"
	"fmt"
)

var x = 0
func main()  {
	done := make(chan bool, 1)
	var wg sync.WaitGroup
	for i:=0; i< 1000; i++ {
		wg.Add(1)
		go func() {
			done <- true
			x = x + 1
			<- done
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x val:",x)
}
