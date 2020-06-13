package main

import (
	"fmt"
	"time"
)

func main() {
	//无缓冲channel
	//两个任务异步执行，然后等待结果返回
	ret1, ret2 := AsyncService(serviceTask), AsyncService(otherTask)
	str1 := <-ret1
	str2 := <-ret2
	fmt.Println("str1 async return:", str1)
	fmt.Println("str2 async return:", str2)
}

func serviceTask() string {
	fmt.Println("service task init...")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("service task over...")
	return "return in service task"
}

func otherTask() string {
	fmt.Println("other task init...")
	time.Sleep(20 * time.Millisecond)
	fmt.Println("other task over...")
	return "return in other task"
}

func AsyncService(fun func() string) chan string {
	//异步执行任务
	ret := make(chan string)
	go func() {
		str := fun()
		ret <- str
	}()
	return ret
}
