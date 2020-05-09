package main

import "fmt"

func main() {
    s := []string{"a", "b"}
    for k,v := range s {
        fmt.Println(k,v)
    }
}
