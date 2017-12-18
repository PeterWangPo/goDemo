package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	start := now.Unix()
	fmt.Println(start)
	for i := 0; i < 100000000; i++ {

	}
	end := time.Now()
	end2 := end.Unix()
	fmt.Println(end2)
	fmt.Println(end2 - start)
}
