package main

import (
	"fmt"
	"sync"
)

var xx = 0

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incre(&wg)
	}
	wg.Wait()
	fmt.Println("x val:", xx)
}

func incre(wg *sync.WaitGroup) {
	xx = xx + 1
	wg.Done()
}
