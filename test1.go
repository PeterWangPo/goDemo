package main

import "fmt"

func main() {
	a := inc()
	fmt.Println(a)
    s := []int{1,2,3}
    fmt.Println(intc(s))
    fmt.Println(s)
}
func inc() (v int) {
	defer func() { v++ }()
	return 42
}
func intc(s []int) ([]int) {
    for k,_ := range s {
        s[k] *= 2
    }
    return s
}
