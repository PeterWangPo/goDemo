package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	//done := make(chan bool)
	done := make(chan struct{})
	go createJobs(100, jobs)
	go createWorker(10, jobs, result)
	go handleResult(done, result)
	<-done
	fmt.Println("all done")
}

//func handleResult(done chan<- bool, result <-chan int) {
func handleResult(done chan<- struct{}, result <-chan int) {
	for r := range result {
		fmt.Println("handle result:", r)
	}
	//done <- true
	done <- struct{}{}
}
func createJobs(jobNum int, jobs chan<- int) {
	for i := 0; i < jobNum; i++ {
		jobs <- i
	}
	close(jobs)
}
func createWorker(workerNum int, jobs <-chan int, result chan<- int) {
	var wg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go toWork(&wg, jobs, result)
	}
	wg.Wait()
	close(result)
}

func toWork(wg *sync.WaitGroup, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		goWork(job, result)
	}
	wg.Done()
}
func goWork(job int, result chan<- int) {
	fmt.Println("do... job:", job)
	result <- job
}
