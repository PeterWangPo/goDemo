package main

import (
    "fmt"
    "context"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    //启动多个子goroutine，看能否全部退出
    for i:=0 ; i< 3; i++ {
        //绑定一个变量后，会生成一个新的context，那如果有多个变量怎么办？递归生成多个，显然不现实？
        c := context.WithValue(ctx,"n",i)
        go work(c)
    }
    time.Sleep(600 * time.Millisecond)
    cancel()
    time.Sleep(600 * time.Millisecond)
    fmt.Println("all done...")
}

func work(ctx context.Context) {
    //子goroutine里面在创建出子goroutine，看是否也会退出
    n := ctx.Value("n")
    go func(n interface{}){
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
