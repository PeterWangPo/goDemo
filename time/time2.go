package main

import (
	"fmt"
	"time"
)

func main() {
	l, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Date(2017, 6, 8, 12, 12, 12, 0, l)
	fmt.Println(t.Location())
	fmt.Println(t.Before(time.Now()))
	fmt.Println(t.Add(3000000))
}
