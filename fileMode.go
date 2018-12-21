package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("fmtTest.go")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fi.IsDir())
}
