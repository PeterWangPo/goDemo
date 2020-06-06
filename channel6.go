package main

import (
	"fmt"
	"time"
)

func main() {
	//两个任务异步执行，谁先返回用谁
	select {
	case str1 := <-AsyncService(serviceTask):
		fmt.Println(str1)
	case str := <-AsyncService(otherTask):
		fmt.Println(str)
	case <-time.After(time.Millisecond * 10):
		fmt.Println("time out")
	}
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
