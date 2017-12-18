package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	url := []byte("www.baidu.com+")
	fmt.Println(url)
	b := base64.URLEncoding.EncodeToString(url)
	fmt.Println(b)
	org, _ := base64.URLEncoding.DecodeString(b)
	fmt.Println(string(org))
}
