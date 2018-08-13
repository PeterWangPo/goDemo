package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Format("01"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix())//当前时间戳
	fmt.Println(time.Unix(1527955046,0).Format("2006-01-02 15:04:05"))//时间戳格式化
	unix := time.Date(2006,01,2,15,4,5,0, time.Local).Unix()//时间格式转换为时间戳
	fmt.Println(unix,time.Unix(unix, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().UnixNano())//返回毫秒戳
	fmt.Println(time.Unix(1527955046,0).Equal(time.Unix(1527955046,1)))//比较两个时间是否相等
	fmt.Println(time.Unix(1527955046,0).Before(time.Unix(1527955046,1)))//比较一个时间是否在另一个前面
	fmt.Println(time.Unix(1527955046,0).After(time.Unix(1527955046,1)))
	fmt.Println(time.Now().Date())//返回年月日
	fmt.Println(time.Now().Clock())//返回时分秒
	fmt.Println(time.Now().YearDay())//返回一年中的第几天
	fmt.Println(time.Now().Day())//返回一月中的第几天
	fmt.Println(time.Now().Weekday())//返回一周中的第几天
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	nowT := now.Add(10000000000)//1秒等于1000000000纳秒
	fmt.Println(nowT.Format("2006-01-02 15:04:05"))
	fmt.Println(now.AddDate(1,1,1).Format("2006-01-02 15:04:05"))
	fmt.Println(now.String())
	fmt.Println(now.MarshalJSON())
}
