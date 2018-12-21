package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(resultFunc(consumerFunc(s, producerFunc(s))))
}
func producerFunc(s []int) <-chan int {
	producer := make(chan int, len(s))
	defer close(producer)
	for _, v := range s {
		producer <- v
	}
	return producer
}
func consumerFunc(s []int, producer <-chan int) <-chan int {
	consumer := make(chan int, len(s))
	defer close(consumer)
	for val := range producer {
		consumer <- int(math.Pow(float64(val), 2))
	}
	return consumer
}
func resultFunc(consumer <-chan int) int {
	sum := 0
	for val := range consumer {
		sum += val
	}
	return sum
}
