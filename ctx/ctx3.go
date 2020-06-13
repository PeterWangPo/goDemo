package main

import (
    "fmt"
    "context"
)

func main() {
    //生成一个空的context
    parent := context.Background()
    //创建一个子context，然后将这个子context当作参数传递给goroutine使用
    ctx,cancel := context.WithCancel(parent)
    runtimeT := 0
    done := make(chan struct{})
    //go func(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("goroutine done")
                done<-struct{}{}
                return
            default:
                fmt.Println("goroutine:", runtimeT)
                runtimeT = runtimeT + 1
            }
            if runtimeT > 5 {
                cancel()
            }
        }
    }()
    //}(ctx)
    <-done
    fmt.Println("all done...")
}
