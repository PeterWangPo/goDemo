package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("%q", line)
	line2, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Printf("%q", line)
	fmt.Println()
	fmt.Printf("%q", line2)
}
