package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	for k, v := range args {
        /*
		if k == 0 {
			continue
		}
        */
		fmt.Println(k,v)
	}
	var num [5][0]int
	for range num {
		fmt.Println("hello world")
	}
}
