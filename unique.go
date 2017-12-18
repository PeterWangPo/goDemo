package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("error %v", err)
	}
	input := bufio.NewScanner(f)
	var s []string
	m := make(map[string]int)
	for input.Scan() {
		m[input.Text()]++
		if m[input.Text()] >= 2 {
			s = append(s, input.Text())
		}
	}
	// fmt.Println(s)
	f2, err2 := os.OpenFile("./b.txt", os.O_APPEND, 0666)
	if err2 != nil {
		fmt.Println("err2 %v", err2)
	}
	if len(s) >= 1 {
		for _, v := range s {
			fmt.Println(os.PathSeparator)
			_, err3 := f2.WriteString(v + string(os.PathSeparator))
			if err3 != nil {
				fmt.Println(err3)
			}
		}

	} else {
		fmt.Println("no unique")
	}
}
