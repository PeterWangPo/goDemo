package main

import (
	"fmt"
	"time"
)

//todo...未完成
//需求:生产者不停的生产,直到收到停止信号,然后退出
func main() {
	cl := make(chan struct{})
	done := make(chan struct{})
	go producer(cl, done)
	time.Sleep(time.Second * 1)
	cl <- struct{}{}
	<-done
}

func producer(cl chan struct{}, done chan struct{}) {
	var i int
	for {
		i += 1
		fmt.Println(i)
		//这里不会阻塞,如果cl没有返回，就会返回default里面的值
		if isCancel(cl) {
			fmt.Println("job done...")
			break
		}
	}
	done <- struct{}{}
}

func isCancel(cl chan struct{}) bool {
	select {
	case <-cl:
		return true
	default:
		return false
	}
}
