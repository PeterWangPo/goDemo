package main

import (
    "fmt"
    "sync"
)

func main() {
    //10个鸡蛋，100个人抢
    eggs := make(chan int, 10)

    for i :=0; i < 10; i++ {
        eggs <-i
    }

    var wg sync.WaitGroup
    for i:=0; i < 100; i++ {
        wg.Add(1)
        go func(num int){
            select {
            case e:=<-eggs:
                fmt.Printf("people %d get eggs %d\n", num, e)
            default:
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}
