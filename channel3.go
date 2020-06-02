package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    //通过锁来解决锁竞争问题
    var lk sync.Mutex
    count := 0

    for i:= 0; i < 5000; i++ {
        go func(){
            defer func() {
                lk.Unlock()
            }()
            lk.Lock()
            count++
        }()
    }
    time.Sleep(1 * time.Second)
    fmt.Println("count:", count)
}
