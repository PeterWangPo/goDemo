package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {

runtime.GOMAXPROCS(runtime.NumCPU())

jobs := make(chan int, 100)
results := make(chan int, 100)

	// 招聘3个工人，让他们工作（待命），每个工人会从jobs管道里获得任务，工人干活，把结果放到results管道
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 老板端，实时查看任务结果
	go func(){
		for v1 := range results {
		fmt.Println("val:", v1)
		}
	}()

// 分配任务，共有26w笔任务
	for j := 1; j <= 26000; j++ {
		jobs <- j
	}
	close(jobs)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
	fmt.Println("worker", id, "started  job", j)
	time.Sleep(time.Second)
	fmt.Println("worker", id, "finished job", j)
	results <- j * 2
	}
}