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
        //为了解决上面ctx6.go里面传递多个值的问题，我们可以传递一个map，这样可以向下传递多个值
        m := make(map[string]interface{});
        m["name"] = 'a' + i
        m["num"] = i
        fmt.Println(m)
        c := context.WithValue(ctx,"m",m)
        go work(c)
    }
    time.Sleep(600 * time.Millisecond)
    cancel()
    time.Sleep(600 * time.Millisecond)
    fmt.Println("all done...")
}

func work(ctx context.Context) {
    //子goroutine里面在创建出子goroutine，看是否也会退出
    n := ctx.Value("m")
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
