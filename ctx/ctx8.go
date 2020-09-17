package main

import (
    "fmt"
    "context"
)

func main() {
    gen := func(ctx context.Context) chan int {
        n := 0
        ret := make(chan int)
        go func(){
            for {
                select {
                case <-ctx.Done():
                    return
                case ret<-n:
                    n++
                }
            }
        }()
        return ret
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    for i := range gen(ctx) {
        fmt.Println(i)
        if i == 10 {
            break
        }
    }
}
