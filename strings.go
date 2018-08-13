package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.ContainsAny("我是abc", "ef"))
	fmt.Println(strings.Contains("我是abc", "我"))
	fmt.Println(strings.HasSuffix("我是abc我", "我"))
	fmt.Println(strings.HasPrefix("我是abc", "a"))
	fmt.Println(strings.Join([]string{"1","2"},","))
	fmt.Println(strings.ContainsRune("我是abc",'a'))
	fmt.Println(strings.Split("abcd",""))
	fmt.Println(strings.Index("abcd","e"))
	fmt.Println(strings.Fields("abcd  de    fg"))
	fmt.Println(strings.Compare("a", "ad"))
	fmt.Println(strings.Count("abcada",""))
	fmt.Println(strings.EqualFold("ab","AB"))
	fmt.Println(strings.FieldsFunc("abc", func(r rune) bool {
		if (r == 'a' || r == 'c') {
			return true
		} else {
			return false
		}
	}))
	fmt.Println(strings.IndexAny("abc",""))
	fmt.Println(strings.IndexFunc("abc", func(r rune) bool {
		if (r == 'a') {
			return true
		} else {
			return false
		}
	}))
	read := strings.NewReader("abcdefg")
	fmt.Println(read.Len())
	b,err := read.ReadByte()
	fmt.Println(b,err)
	err  = read.UnreadByte()
	fmt.Println(err)
}
