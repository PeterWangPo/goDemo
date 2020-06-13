package main

import (
    "fmt"
    "context"
    "time"
)

//通过context可以控制基于这个context或者衍生出的子context，优雅的退出
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    //启动多个子goroutine，看能否全部退出
    go work(ctx, 1)
    go work(ctx, 2)
    go work(ctx, 3)
    time.Sleep(600 * time.Millisecond)
    cancel()
    time.Sleep(600 * time.Millisecond)
    fmt.Println("all done...")
}

func work(ctx context.Context, n int32) {
    //子goroutine里面在创建出子goroutine，看是否也会退出
    go func(n int32){
        for{
            select {
            case <-ctx.Done():
                fmt.Printf("sub goroutine %d done\n", n)
                return
            default:
                time.Sleep(50 * time.Millisecond)
                fmt.Printf("sun goroutine %d running\n", n)
            }
        }
    }(n)
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("goroutine %d done\n", n)
            return
        default:
            time.Sleep(200 * time.Millisecond)
            fmt.Printf("goroutine %d running\n", n)
        }
    }
}
