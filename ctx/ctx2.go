package main

import (
    "fmt"
    "time"
)

//通过channel来通信告诉goroutine结束.这也是官方建议的一种方式，但是也存在局限性。
//如果有很多个goroutine都需要控制结束怎么办？案例中用了两个goroutine，就使用了两个channel
//来控制结束，如果goroutine又衍生出了更多个goroutine呢，或者说是一层一层的无穷尽的goroutine呢？
//那像这样定义channel来控制结束显然无法解决这样的问题，因为goroutine的关系链就导致了这种场景非常复杂
func main () {
    done := make(chan struct{})
    done2 := make(chan struct{})

    go func(){
        for {
            select {
            case <-done:
                fmt.Println("goroutine done...")
                done2<-struct{}{}
                return
            default:
                time.Sleep(50 * time.Millisecond)
                fmt.Println("goroutine running...")
            }
        }
    }()
    go func(){
        time.Sleep(time.Millisecond * 200)
        done<-struct{}{}
    }()
    <-done2
    fmt.Println("all done...")
}
