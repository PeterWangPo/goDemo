package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"a": "111", "b": "222"}
	for v, ok := range a {
		fmt.Println(v, ok)
	}
	if v, ok := a["a"]; ok {
		fmt.Println("find", v)
	} else {
		fmt.Println("not find", v)
	}
}
