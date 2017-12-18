package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())
	rand.Seed(int64(time.Now().Nanosecond()))
	r := rand.Intn(15)
	fmt.Println(r)
	n := time.Unix(time.Now().Unix(), 0)
	fmt.Println(n.Format("2006-01-02 15:04:05"))
}
