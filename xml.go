package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Sitemapindex struct {
	XMLName xml.Name  `xml:"sitemapindex"`
	Xlns    string    `xml:"xmlns,attr"`
	Sitemap []Sitemap `xml:"sitemap"`
}
type Sitemap struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func main() {
	var i Sitemapindex
	i.Xlns = "www.baidu.com"
	i.Sitemap = []Sitemap{
		{Loc: "1", Lastmod: "1"},
		{Loc: "2", Lastmod: "2"},
		{Loc: "3", Lastmod: "3"},
		{Loc: "4", Lastmod: "4"},
	}
	if data, err := xml.Marshal(&i); err != nil {
		panic(err)
	} else {
		if f, err1 := os.OpenFile("./sitemap.xml", os.O_CREATE, 0666); err1 != nil {
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
