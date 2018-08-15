package main

import (
	"sync"
	"fmt"
)

var x = 0
func main()  {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incre(&wg)
	}
	wg.Wait()
	fmt.Println("x val:",x)
}

func incre(wg *sync.WaitGroup)  {
	x = x+1
	wg.Done()
}