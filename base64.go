package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	//"io"
	"bufio"
	//"strconv"
	//"path/filepath"
)

const filename = "base64.txt"

func main() {
	url := []byte("www.baidu.com+")
	fmt.Println(url)
	b := base64.URLEncoding.EncodeToString(url)
	fmt.Println(b)
	org, _ := base64.URLEncoding.DecodeString(b)
	fmt.Println(string(org))
	fmt.Println("=================")
	read()
	write()
	write2()
	write3()
}

//流式解码
func read() {
	str := "dGhpcyBpcyBhIGJhc2U2NCBlbmNvZGUgdGVzdCEhICBoYW8tbG92ZS9pdCBpcy4="
	read := strings.NewReader(str)
	ioReader := base64.NewDecoder(base64.StdEncoding, read)
	src := make([]byte, 6)
	target := ""
	for {
		n, err := ioReader.Read(src)
		if err != nil || n == 0 {
			break
		}
		target += string(src[:n])
	}
	fmt.Println("流式解码后的内容:" + target)
}

//流式加密,输出到os.Stdout
func write() {
	str := []byte("this is hhh! a.?!-- test 123")
	writer := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	if _, err := writer.Write(str); err != nil {
		panic(err)
	}
	writer.Close()
	fmt.Println()
}

//流式生成加密字符串后，写入到文件里面
func write2() {
	//坑一：不能用Open方法,Open只读
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 666)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	//生成一个io.Writer对象，可通过下面的方法  写入文件可以传入一个file对象
	w := bufio.NewWriter(file)
	//bufio有缓存，实际是写入buffer,这里需要手动刷缓存
	defer w.Flush()
	str := []byte("this is hhh! a.?!-- test 123")
	//生成一个流加密器
	writer := base64.NewEncoder(base64.StdEncoding, w)
	if _, err := writer.Write(str); err != nil {
		panic(err)
	}
	writer.Close()
	fmt.Println()
}

func write3() {
	fmt.Println("||||||||||||||||||||||||||")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("this is hhh! a.?!-- test 123")))
	encode := base64.NewEncoding(strings.Repeat("-", 64))
	dst := encode.EncodeToString([]byte("this is hhh! a.?!-- test 123"))
	fmt.Println(dst)
	fmt.Println(len(dst)%4 == 0)
	/*
	   var dst []byte
	   encode.Encode(dst, []byte("this is hhh! a.?!-- test 123"))
	   fmt.Println(string(dst))
	*/
}
