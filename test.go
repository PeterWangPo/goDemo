package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//fmt.Println(getRandomString(3))
	//fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(3))
	//fmt.Println(t())
	//fmt.Println(GetlastMonthLastDayUnix())
	file, err := os.OpenFile("./b.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("os openfile err：", err)
	}
	str := "abcdefg"
	//byteStr := []byte(str)
	//n, err := file.Write(byteStr)
	n, err := file.WriteString(str)
	if err != nil {
		fmt.Println("write file err:", err)
		fmt.Println(n)
	}
	fmt.Println("done...")
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func t() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	ctime := time.Date(year, month, 01, 0, 0, 0, 0, time.Local)
	return ctime.AddDate(0, -1, 0).Unix()
	//return ctime.Unix()
}

//获取当前月第一天开始时间戳
func GetCurrentMonthFirstDayUnix() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	return time.Date(year, month, 01, 0, 0, 0, 0, time.Local).Unix()
}

//获取上一个月第一天时间戳
func GetlastMonthFirstDayUnix() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	return time.Date(year, month, 01, 0, 0, 0, 0, time.Local).AddDate(0, -1, 0).Unix()
}

//获取上一个月最后一天时间戳
func GetlastMonthLastDayUnix() int64 {
	return GetCurrentMonthFirstDayUnix() - 1
}
