package main

import (
    "fmt"
    "sync"
)

//并发控制有两种经典的方式，一种是WaitGroup，另外一种是Context
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        fmt.Println("1 done...")
        wg.Done()
    }()
    go func() {
        fmt.Println("2 done...")
        wg.Done()
    }()
    wg.Wait()
}
//以上控制并发的方式，适用于多个goroutine协同做一件事情，因为每个goroutine做的都是这个事情的一部分
//只有全部goroutine都完成了，这件事才算完成。这也是go里面通过共享变量来通信的一个案例，但是go里面一般
//不这么使用，一般是通过channel来通信
