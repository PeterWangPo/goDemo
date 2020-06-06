package main

import (
	"fmt"
	"time"
)

func main() {
	//锁竞争
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count:", count)
}
