package main

import (
	"fmt"
	"os"
)

func main() {
	//Hostname返回内核提供的主机名
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("name is %s", name)
}
