package main

import (
	"fmt"
)

func main() {
	i := 0
here:
	fmt.Println(i)
	i++
	goto here
}
