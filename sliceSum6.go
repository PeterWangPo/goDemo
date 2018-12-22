package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(resultFunc3(consumerFunc3(s, producerFunc3(s))))
}
func producerFunc3(s []int) <-chan int {
	producer := make(chan int, len(s))
	defer close(producer)
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	for _, v := range s {
		go func(val int, inCh chan<- int, group *sync.WaitGroup) {
			inCh <- val
			group.Done()
		}(v, producer, &wg)
	}
	wg.Wait()
	return producer
}

func consumerFunc3(s []int, producer <-chan int) <-chan int {
	consumer := make(chan int, len(s))
	defer close(consumer)
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	for val := range producer {
		go func(val int, outCh chan<- int, group *sync.WaitGroup) {
			outCh <- int(math.Pow(float64(val), 2))
			group.Done()
		}(val, consumer, &wg)
	}
	wg.Wait()
	return consumer
}

func resultFunc3(consumer <-chan int) int {
	sum := 0
	for val := range consumer {
		sum += val
	}
	return sum
}
