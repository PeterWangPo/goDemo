package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(resultFunc2(consumerFunc2(s, producerFunc2(s))))
}
func producerFunc2(s []int) <-chan int {
	producer := make(chan int, len(s))
	defer close(producer)
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	for _, v := range s {
		go handleProducer(producer, v, &wg)
	}
	wg.Wait()
	return producer
}
func handleProducer(producer chan<- int, val int, group *sync.WaitGroup) {
	producer <- val
	group.Done()
}
func consumerFunc2(s []int, producer <-chan int) <-chan int {
	consumer := make(chan int, len(s))
	defer close(consumer)
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	for val := range producer {
		go handleConsumer(consumer, val, &wg)
	}
	wg.Wait()
	return consumer
}
func handleConsumer(consumer chan<- int, val int, group *sync.WaitGroup) {
	consumer <- int(math.Pow(float64(val), 2))
	group.Done()
}
func resultFunc2(consumer <-chan int) int {
	sum := 0
	for val := range consumer {
		sum += val
	}
	return sum
}
