package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, len(s))
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	sum := 0
	go func() {
		wg.Wait()
		close(result)
	}()
	for _, v := range s {
		go getSum3(v, result, &wg)
	}
	for val := range result {
		sum += val
		fmt.Println("total item val:", sum)
	}
	fmt.Println(sum)
}

func getSum3(val int, result chan<- int, group *sync.WaitGroup) {
	result <- int(math.Pow(float64(val), 2))
	group.Done()
}
