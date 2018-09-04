package main

import (
	"net/url"
	"log"
	"fmt"
)

func main()  {
	u := "https://detail.tmall.com/auction/noitem.htm?type=2"
	uq := url.QueryEscape(u)
	fmt.Println(uq)
	fmt.Println(url.QueryUnescape(uq))
	us , err := url.Parse(u)
	if err != nil {
		log.Fatal("parse url err:",err)
	}
	id := us.Query().Get("id")
	fmt.Println("id is :", id)
}
