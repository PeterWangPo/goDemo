package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间戳
	now := time.Now()
	fmt.Println(now.Unix())
	//当前时间
	fmt.Println(now.Format("2006-01-02 15:04:05")) //必须是这个时间点，这个时间是golang的诞生时间
	fmt.Println(now.Format("15:04:05"))
	//指定时间戳转换为时间
	t := 1495866374
	fmt.Println(time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"))
	//指定时间转化为时间戳
	fmt.Println(time.Date(2017, 5, 27, 14, 26, 14, 0, time.Local).Unix())
	fmt.Println(now.Second())
}
