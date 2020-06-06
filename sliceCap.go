package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5}
	sli := s[2:4]
	fmt.Println("len:", len(sli), "cap:", cap(sli))
	sli = append(sli, 6)
	fmt.Println("len:", len(sli), "cap:", cap(sli))
	sli = append(sli, 7)
	fmt.Println("len:", len(sli), "cap:", cap(sli))
}
