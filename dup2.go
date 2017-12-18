package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	counts := make(map[string]int)
	for _, v := range args {
		f, err := os.Open(v)
		if err != nil {
			fmt.Printf("error %v", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			//			fmt.Println(input.Text())
			counts[input.Text()]++
		}
		f.Close()
	}
	fmt.Println(counts)
}
