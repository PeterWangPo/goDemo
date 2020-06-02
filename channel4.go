package main

import (
    "fmt"
    "sync"
)

func main() {
    //通过同步队列来确保协程执行完毕
    var lk sync.Mutex
    var wg sync.WaitGroup
    count := 0
    for i := 0; i < 5000; i++ {
        wg.Add(1)
        go func(){
            defer func(){
                lk.Unlock()
            }()
            lk.Lock()
            count++
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("count:", count)
}
