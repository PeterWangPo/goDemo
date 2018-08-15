package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			x = x + 1
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x val:", x)
}
