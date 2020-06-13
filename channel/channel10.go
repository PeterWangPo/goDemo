package main

import (
    "fmt"
    "time"
)

func main() {
	//0.要通过通信来共享内存，而不是通过共享内存来通信
	//1.向一个nil channel发送消息，会一直阻塞
	//2.向一个已经关闭的channel发送消息，会panic
	//3.channel关闭后不可以继续向channel发送消息，但是可以继续从channel接受消息
	//4.channel只能有发送数据方关闭
	//5.当channel关闭且缓冲区为空时，继续从channel获取数据，得到的时对应类型的零值
	//6.for i := range ch 通过for range可以从channel中接受数据，如果ch没有被关闭，那么for range 会一直阻塞，当channel关闭后，for range 会退出循环
	//7.无缓冲channel：接受者会阻塞直到接收到消息，发送者会阻塞直到接收者接收到消息，这种channel可以用于两个goroutine进行状态同步
	//8.有缓冲channel：当缓冲区已满时，发送者会阻塞，接收者不会阻塞；当缓冲区没有满，切不为空时，两者都不阻塞；当缓冲区为空时，接受者阻塞
    //9.需要close的地方一定要close,否则会有一些意想不到的bug，常见需要close的地方：
    //9.1 file open/close
    //9.2 lock lock/unlock
    //9.3 db查询 查询出db数据后要close
    //9.4 channl 发送者发送完毕后需要close
    //9.5 ....其他未知地方

    //常见deadlock:

    //1.无缓冲channel, 只有接收者和发送者都准备好，才不会死锁,如下会死锁
    //下面看上去正常，一个发，一个收，但是发的时候，收还没有准备好，所以会阻塞，导致死锁
    //=======
    //ch := make(chan string)
    //ch <- "hello world"
    //fmt.Println(<-ch)
    // ============

    //上面如下修改可正常：1.让发或者收异步准备好 2.加入缓存
    //ch := make(chan string)
    //go func(){
    //    ch <- "hello world"
    //}()
    //time.Sleep(time.Millisecond)
    //fmt.Println(<-ch)

    //使用缓冲,容量为1
    //time.Sleep(time.Millisecond)
    //ch := make(chan string, 1)
    //ch <-"hello world"
    //fmt.Println(<-ch)

    //2.for range 读channel里面的数据，发送者不关闭会阻塞，导致deadlock

    //3.写入数据量，大于缓冲容量，就会死锁
    //time.Sleep(time.Millisecond)
    //ch := make(chan string, 1)
    //ch <-"e"
    //ch <-"b"
    //fmt.Println(<-ch, <-ch)

    //如果有异步消耗，则不会出现问题
    //ch := make(chan string, 1)
    //go func(){
    //    //for i := range ch {
    //    //    fmt.Println(i)
    //    //}
    //    fmt.Println(<-ch)
    //}()
    //ch <-"a"
    //ch <-"c"
    //time.Sleep(time.Millisecond)
}
