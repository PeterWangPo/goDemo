package main

import (
    "fmt"
    "strings"
)

func main() {
    s1 := "abc sdfsfaffabc sdfab123badabcabcbadabc"
    fmt.Println(rp(s1, "abc", "ABC", -1))
}
func rp(target, old, n string) string {
    return strings.Replace(target, old, new, -1)
}

