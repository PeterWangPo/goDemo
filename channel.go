package main

import (
	"fmt"
	//"runtime"
	//"time"
)

func main() {

	//runtime.GOMAXPROCS(runtime.NumCPU())

	jobs := make(chan int, 100)
	results := make(chan int, 100)
    done := make(chan struct{}, 1)

	// 招聘3个工人，让他们工作（待命），每个工人会从jobs管道里获得任务，工人干活，把结果放到results管道
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 老板端，实时查看任务结果
	go func() {
		for v1 := range results {
			fmt.Println("val:", v1)
		}
        done <- struct{}{}
	}()

	// 分配任务，共有26w笔任务
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
    close(jobs)
    <-done
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		results <- j * 2
	}
    //此处close有问题，因为如果一直一个work检测到jobs关闭了，执行完成后就会关闭results，如果另外的goroutine再后面执行完，然后写入数据到results，就会报错!
    close(results)
}
