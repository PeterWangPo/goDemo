package main

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
}
