package main

import (
	"fmt"
	"time"
)

//todo...未完成
//需求:生产者不停的生产,直到收到停止信号,然后退出
func main() {
	var task []int
	tm := time.Tick(time.Second)
	tf := time.After(time.Second * 4)
	for {
		select {
		case <-tf:
			fmt.Println("task execute 2 seconds...")
			return
		case <-tm:
			fmt.Println("task len:", len(task))
		default:
			task = append(task, 1)
		}
	}
}
