package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	// "strconv"
	"strings"
	"time"
)

const _URL_PREFIX = "http://download.miyabaobei.com/xml/360_feed/"

const _IMG_HOST = "https://img07.miyabaobei.com/"

//for test env
const _DB = "write_user:write_pwd@tcp(172.16.104.207:3307)/mia_test2?charset=utf8"

//价格单位
const _PRICE_UNIT = "RMB"

//是否可用
const _AVAILABILITY = 1

//卖家官方网址
const _SELLER_SITE_URL = "https://www.mia.com"

//卖家官方名称
const _SHOP_NAME = "蜜芽宝贝"

const _SELLER_NAME = "蜜芽宝贝"

const _LOGO = "http://img03.miyabaobei.com/d1/p5/2017/08/18/41/ff/41ff9ac22688a1f80af29ef7f57a41e3369817830.jpg"

func main() {
	type Images struct {
		//内部属性都要大写字母开头，属性节点的名称变量名固定为XMLName，内部的文本统一叫innerxml
		Index     int    `xml:"index,attr"` //表示属性
		InnerText string `xml:",innerxml"`  //表示文本
	}
	type Mimages struct {
		XMLName xml.Name `xml:"moreImages"`
		Img     []Images `xml:"img"`
	}
	type SubAttribute struct {
		Key   string `xml:"key"`
		Value string `xml:"value"`
	}

	type Attribute struct {
		Attribute []SubAttribute `xml:"attribute"`
	}
	type Data struct {
		OuterID string  `xml:"outerID"`
		Name    string  `xml:"name"`
		Price   float64 `xml:"price"`
		Value   float64 `xml:"value"`
		// Saving           float64      `xml:"saving"`
		PriceUnit        string    `xml:"priceUnit"`
		Availability     int       `xml:"availability"`
		Image            string    `xml:"image"`
		Brand            string    `xml:"brand"`
		Loc              string    `xml:"loc"`
		PcLoc            string    `xml:"pcLoc"`
		SellerSiteUrl    string    `xml:"sellerSiteUrl"`
		ShopName         string    `xml:"shopName"`
		SearchWiseUrl    string    `xml:"searchWiseUrl"`
		Category         string    `xml:"category"`
		CategoryUrl      string    `xml:"categoryUrl"`
		CategoryPcUrl    string    `xml:"categoryPcUrl"`
		SubCategory      string    `xml:"subCategory"`
		SubCategoryUrl   string    `xml:"subCategoryUrl"`
		SubcategoryPcUrl string    `xml:"subcategoryPcUrl"`
		SellerName       string    `xml:"sellerName"`
		Logo             string    `xml:"logo"`
		MoreImages       Mimages   `xml:"moreImages"`
		Choice           Attribute `xml:"choice"`
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
	pageSize := "300"
	maxId := "0"
	//初始化sitemapindex数据
	var sitemapindex Sitemapindex
	var sitemap []Sitemap
	for {
		sql := "select a.id as outerID,a.name,a.sale_price as price, a.category_id,a.market_price as value,b.name as brand from item as a left join item_brand as b on a.brand_id = b.id where a.status = 1 and a.id >" + maxId + " order by a.id asc limit 0," + pageSize
		sql2 := "select count(*) as total from item as a left join item_brand as b on a.brand_id = b.id where a.status = 1 and a.id >" + maxId
		var total int
		if rowCountErr := db.QueryRow(sql2).Scan(&total); rowCountErr != nil {
			panic(rowCountErr)
		} else {
			fmt.Println("有效数据条数：", total)
			//查询总数
			//总数大于零就执行查询，否则就跳出循环
			if total > 0 {
				var r Urlset   //初始化xml
				var urls []Url //初始化urls
				//查询具体的数据
				if rows, rowsErr := db.Query(sql); rowsErr != nil {
					panic(rowsErr)
				} else {
					for rows.Next() {
						var (
							outerID, name, brand, category_id string
							price, value                      float64
						)
						rows.Scan(&outerID, &name, &price, &category_id, &value, &brand)
						if outerID == "" {
							continue
						}
						maxId = outerID
						//loc获取
						_loc := getLocUrl(outerID)
						//通过categoryId获取分类信息
						_categoryInfo := getCategoryInfoByParentId(category_id, db)
						// fmt.Println(_categoryInfo)
						// os.Exit(1)
						if _categoryInfo["categoryId"] == "" {
							continue
						}
						//获取商品图片
						_imgs := getImages(outerID, db)
						if len(_imgs) <= 0 {
							continue
						}
						_firstImg := ""
						var _moreImages []Images
						for i, v := range _imgs {
							if i == 0 {
								_firstImg = v
							} else {
								_moreImages = append(_moreImages, Images{InnerText: v, Index: i})
							}
						}
						moreImages := Mimages{Img: _moreImages}
						// moreImages := _moreImages
						subAttribute := []SubAttribute{SubAttribute{Key: "ext_down_load", Value: "https://m.mia.com/detail-a-" + outerID + ".html"}, SubAttribute{Key: "ext_put_url", Value: "https://m.mia.com/detail-a-" + outerID + ".html"}}
						//Data数据
						_choice := Attribute{Attribute: subAttribute}
						itemData := Data{
							OuterID:          outerID,
							Name:             name,
							Price:            price,
							Value:            value,
							PriceUnit:        _PRICE_UNIT,
							Availability:     _AVAILABILITY,
							Image:            _firstImg,
							MoreImages:       moreImages,
							Brand:            brand,
							Loc:              _loc,
							PcLoc:            getPcLoc(outerID),
							SellerSiteUrl:    _SELLER_SITE_URL,
							ShopName:         _SHOP_NAME,
							SearchWiseUrl:    getSearchWiseUrl(outerID),
							Category:         category_id,
							CategoryUrl:      getCategoryUrl(category_id),
							CategoryPcUrl:    getCategoryPcUrl(category_id),
							SubCategory:      _categoryInfo["subCategoryId"],
							SubCategoryUrl:   getCategoryUrl(_categoryInfo["subCategoryId"]),
							SubcategoryPcUrl: getCategoryPcUrl(_categoryInfo["subCategoryId"]),
							Choice:           _choice,
							SellerName:       _SELLER_NAME,
							Logo:             _LOGO}
						// fmt.Println(itemData)
						// os.Exit(1)
						//url数据
						itemUrl := Url{Loc: _loc, Data: itemData}
						//urls数据
						urls = append(urls, itemUrl)
						//记录每次最大的outerID

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

func getChoices() {

}

func getImages(itemId string, db *sql.DB) []string {
	imgs := []string{}
	sql := "select local_url from item_pictures where `item_id` = " + itemId + " and `status` = 1 and `type` = 'topic' order by `index` asc limit 4"
	if rows, rowErr := db.Query(sql); rowErr != nil {
		fmt.Println("select img err", rowErr)
		os.Exit(1)
		return imgs
	} else {
		var local_url string
		for rows.Next() {
			rows.Scan(&local_url)
			if local_url != "" {
				imgs = append(imgs, _IMG_HOST+local_url)
			}
		}
	}
	return imgs
}

func generateXml(filename string, data interface{}, callback func()) {
	if xmlData, err4 := xml.Marshal(&data); err4 != nil {
		panic(err4)
	} else {
		if f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666); err != nil {
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

func getCategoryInfoByParentId(parentCategoryId string, db *sql.DB) map[string]string {
	res := make(map[string]string, 4)
	res["categoryId"] = ""
	res["subCategoryId"] = ""
	res["categoryName"] = ""
	res["subCategoryName"] = ""
	sql := "select path from item_category where id = " + parentCategoryId + " and status = 1"
	if rows, rowErr := db.Query(sql); rowErr != nil {
		fmt.Println("select category err", rowErr)
		return res
	} else {
		var path string
		for rows.Next() {
			rows.Scan(&path)
		}
		_categoryIds_ := strings.Split(path, "-")
		_count := len(_categoryIds_)
		if _count < 2 {
			return res
		}
		_parentId := _categoryIds_[0]
		_categoryIds := strings.Replace(path, "-", ",", -1)
		sql2 := "select id,name from item_category where id in (" + _categoryIds + ") and status = 1 order by id asc"
		if subRow, subRowErr := db.Query(sql2); subRowErr != nil {
			fmt.Println("sub category search err")
			return res
		} else {
			for subRow.Next() {
				var id, name string
				subRow.Scan(&id, &name)
				if id == _parentId {
					res["categoryId"] = id
					res["categoryName"] = name
				} else {
					res["subCategoryId"] = id
					res["subCategoryName"] = name
				}
			}
		}
		return res
	}
}

//获取xml文件名
func getXmlName() string {
	return time.Now().Format("20060102150405") + ".xml"
}

//获取当前时间
func getCurrentTime() string {
	return time.Now().Format("2006-01-06 15:04:05")
}

//获取wap对应的详情页
func getLocUrl(itemId string) string {
	return "https://m.mia.com/item-" + itemId + ".html"
}

//获取pc对应的详情页
func getPcLoc(itemId string) string {
	return "https://www.mia.com/item-" + itemId + ".html"
}

//获取wap搜索url
func getSearchWiseUrl(itemId string) string {
	return "https://m.mia.com/detail-a-" + itemId + ".html"
}

//获取wap分类url
func getCategoryUrl(categoryId string) string {
	if categoryId == "" {
		return ""
	}
	return "https://m.mia.com/s/cat" + categoryId + "_jh1.html"
}

//获取pc端的分类url
func getCategoryPcUrl(categoryId string) string {
	if categoryId == "" {
		return ""
	}
	return "https://www.mia.com/search/s?cat=" + categoryId
}
