package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	//如果是函数内部，则是类型是按照定义的顺序加载。。在函数外，全局加载，在函数内使用的话，无需考虑顺序
	//和xml.go的区别：Sitemap这个类型必须定义在Sitemapindex类型前面
	type Sitemap struct {
		Loc     string `xml:"loc"`
		Lastmod string `xml:"lastmod"`
	}
	type Sitemapindex struct {
		XMLName xml.Name  `xml:"sitemapindex"`
		Xmlns   string    `xml:"xmlns,attr"`
		Sitemap []Sitemap `xml:"sitemap"`
	}

	var result Sitemapindex
	result.Xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	result.Sitemap = []Sitemap{
		{Loc: "1", Lastmod: "1"},
		{Loc: "2", Lastmod: "2"},
		{Loc: "3", Lastmod: "3"},
	}
	if data, err := xml.Marshal(&result); err != nil {
		panic(err)
	} else {
		if f, err1 := os.OpenFile("./360.xml", os.O_CREATE, 0666); err1 != nil {
			panic(err1)
		} else {
			if _, err2 := f.WriteString(string(data)); err2 != nil {
				panic(err2)
			} else {
				fmt.Println("done...")
			}
		}
	}
}
