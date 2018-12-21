package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, len(s))
	sum := 0
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(s))
		for _, v := range s {
			go getSum2(v, result, &wg)
		}
		wg.Wait()
		close(result)
	}()

	for val := range result {
		sum += val
		fmt.Println("total item val:", sum)
	}
	fmt.Println(sum)
}

func getSum2(val int, sum chan<- int, group *sync.WaitGroup) {
	sum <- int(math.Pow(float64(val), 2))
	group.Done()
}
