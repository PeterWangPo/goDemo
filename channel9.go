package main

import (
    "fmt"
    "time"
    //"sync"
)

//需求:生产者不停的生产,直到收到停止信号,然后退出;消费者继续消费，直到消费完毕
func main() {
    cl := make(chan struct{})
    cs := make(chan int, 10000)
    done := make(chan struct{})
    go producer(cl, cs)
    go consumer(cs, done)
    time.Sleep(time.Second * 1)
    cl <-struct{}{}
    <- done
    fmt.Println("all done...")
}

func consumer(ch chan int, done chan struct{}) {
    for v := range ch {
        fmt.Println("consumer v:", v)
    }
    done <- struct{}{}
}

func producer(cl chan struct{}, cs chan int) {
    var i int
    for {
        i +=1
        fmt.Println("producer i:", i)
        cs <-i
        //这里不会阻塞,如果cl没有返回，就会返回default里面的值
        if isCancel(cl) {
            break
        }
    }
    close(cs)
}

func isCancel(cl chan struct{}) bool {
    select {
    case <-cl:
        return true
    default:
        return false
    }
}
