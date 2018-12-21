package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, len(s))
	sum := 0
	for _, v := range s {
		go getSum(v, result)
	}
	for i := 0; i < len(s); i++ {
		sum += <-result
		fmt.Println("total item val:", sum)
	}
	fmt.Println(sum)
}

func getSum(val int, sum chan<- int) {
	sum <- int(math.Pow(float64(val), 2))
}
