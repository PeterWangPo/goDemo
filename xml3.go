package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

const _URL_PREFIX = "http://download.miyabaobei.com/xml/360_feed/"
const _DB = "write_user:write_pwd@tcp(172.16.104.207:3307)/mia_test2?charset=utf8"

func main() {
	var getXmlName = func() string {
		return time.Now().Format("20060102150405") + ".xml"
	}
	var getCurrentTime = func() string {
		return time.Now().Format("2006-01-06 15:04:05")
	}
	type Data struct {
		OuterID    string  `xml:"outerID"`
		Name       string  `xml:"name"`
		Price      float64 `xml:"price"`
		Value      float64 `xml:"value"`
		Saving     float64 `xml:"saving"`
		Image      string  `xml:"image"`
		Brand      string  `xml:"brand"`
		TargetUrl  string  `xml:"targetUrl"`
		SellerName string  `xml:"sellerName"`
		Logo       string  `xml:"logo"`
	}
	type Url struct {
		Loc  string `xml:"loc"`
		Data Data   `xml:"data"`
	}
	type Urlset struct {
		XMLName xml.Name `xml:"urlset"`
		Url     []Url    `xml:"url"`
	}
	//sitemap.xml
	type Sitemap struct {
		Loc     string `xml:"loc"`
		Lastmod string `xml:lastmod`
	}
	type Sitemapindex struct {
		XMLName xml.Name  `xml:"sitemapindex"`
		Sitemap []Sitemap `xml:"sitemap"`
	}
	db, dbErr := sql.Open("mysql", _DB)
	if dbErr != nil {
		fmt.Println("db connection err")
		os.Exit(1)
	} else {
		fmt.Println("db ok")
	}
	//初始化sql数据
	pageSize := "5000"
	maxId := "0"
	//初始化sitemapindex数据
	var sitemapindex Sitemapindex
	var sitemap []Sitemap
	for {
		sql := "select a.id as outerID,a.name,a.sale_price as price,a.market_price as value,b.name as brand from item as a left join item_brand as b on a.brand_id = b.id where a.status = 1 and a.id >" + maxId + " order by a.id asc limit 0," + pageSize
		sql2 := "select count(*) as total from item as a left join item_brand as b on a.brand_id = b.id where a.status = 1 and a.id >" + maxId + " order by a.id asc"
		if rowsCount, rowCountErr := db.Query(sql2); rowCountErr != nil {
			panic(rowCountErr)
		} else {
			// fmt.Println(sql)
			//查询总数
			for rowsCount.Next() {
				var total int
				rowsCount.Scan(&total)
				fmt.Println(total)
				//总数大于零就执行查询，否则就跳出循环
				if total > 0 {
					var r Urlset           //初始化xml
					urls := make([]Url, 0) //初始化urls
					//查询具体的数据
					if rows, rowsErr := db.Query(sql); rowsErr != nil {
						panic(rowsErr)
					} else {
						for rows.Next() {
							var (
								outerID, name, image, brand, targetUrl, sellerName, logo string
								price, value, saving                                     float64
							)
							var loc = "www.baidu.com"
							rows.Scan(&outerID, &name, &price, &value, &brand)
							saving = value - price
							image = "www.baidu.com"
							targetUrl = "www.baidu.com"
							logo = "www.baidu.com"
							sellerName = "abc"
							//Data数据
							itemData := Data{OuterID: outerID, Name: name, Price: price, Value: value, Saving: saving, Image: image, Brand: brand, TargetUrl: targetUrl, SellerName: sellerName, Logo: logo}
							//url数据
							itemUrl := Url{Loc: loc, Data: itemData}
							//urls数据
							urls = append(urls, itemUrl)
							//记录每次最大的outerID
							maxId = outerID
						}
						//生成xml数据
						r.Url = urls
					}
					//生成xml文件
					if len(urls) > 1 {
						//取数据,
						fileName := getXmlName()
						generateXml(fileName, r, func() {
							xmlUrl := _URL_PREFIX + fileName
							itemSiteMap := Sitemap{Loc: xmlUrl, Lastmod: getCurrentTime()}
							sitemap = append(sitemap, itemSiteMap)
							//需要停顿几秒，待写文件完毕。主要是防止最后一次写文件没写完，程序就执行完退出
							time.Sleep(5)
							fmt.Println("done...")
						})
					}
				} else {
					goto here
				}
			}

		}

	}
here:
	if len(sitemap) > 0 {
		sitemapindex.Sitemap = sitemap
		generateXml("./sitemap.xml", sitemapindex, func() {
			time.Sleep(2)
			fmt.Println("sitemap done...")
		})
	}
	fmt.Println("all done...")
}
func generateXml(filename string, data interface{}, callback func()) {
	if xmlData, err4 := xml.Marshal(&data); err4 != nil {
		panic(err4)
	} else {
		if f, err := os.OpenFile(filename, os.O_CREATE, 0666); err != nil {
			panic(err)
		} else {
			if _, err2 := f.WriteString(string(xmlData)); err2 != nil {
				panic(err2)
			} else {
				callback()
			}
		}
	}
}
