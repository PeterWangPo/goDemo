package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_URL                 = "https://adx.toutiao.com/adx/retargeting-portal/api?method=test&type={type}&source={source}"
	_TOKEN               = "9d74b3798d934d49adbff4c88a9f3450"
	_ADVERTISER_ID       = 58980835622
	_ADVERTISER_NAME     = "蜜芽宝贝"
	_SOURCE              = "30000001"
	_TABLE_ITEM_PIC      = "item_pictures"
	_TABLE_CART_V4       = "cart_v4"
	_TABLE_ITEM          = "item"
	_TABLE_ITEM_BRAND    = "item_brand"
	_TABLE_ITEM_CATEGORY = "item_category"
)

var _PLATFORM = [...]string{"Android", "IOS"}
var _TYPE = [...]string{"1", "2"}
var _BEHAVIOR_CODE = [...]int{1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 2001, 2002, 2003, 2004, 2005, 2006, 2007}

//post json struct
type PostData struct {
	Items []Item `json:"data"`
}
type Item struct {
	Behavior_code int            `json:"behavior_code"`
	Platform      string         `json:"platform"`
	Mid           string         `json:"mid"`
	Timestamp     int64          `json:"timestamp"`
	Product_infos []Product_info `json:"product_infos"`
}
type Product_info struct {
	Product_id  int     `json:"product_id"`
	Price       float64 `json:"price"`
	Brand_id    string  `json:"brand_id"`
	Category_id string  `json:"category_id"`
}

func main() {
	requestUrl := _getSignatureUrl(_TYPE[0], _SOURCE)
	var data PostData
	var items []Item
	var product_infos []Product_info
	product_infos = append(product_infos, Product_info{Product_id: 1000984, Price: 12.1, Brand_id: "38", Category_id: "123"})
	product_infos = append(product_infos, Product_info{Product_id: 1000985, Price: 16, Brand_id: "32", Category_id: "131"})
	items = append(items, Item{Behavior_code: 2003, Platform: "IOS", Mid: "CC9A40D3-D9A0-4E58-A62F-60DAA626F0A5", Timestamp: time.Now().Unix(), Product_infos: product_infos})
	items = append(items, Item{Behavior_code: 2003, Platform: "IOS", Mid: "CC9A40D3-D9A0-4E58-A62F-60DAA626F0A7", Timestamp: time.Now().Unix(), Product_infos: product_infos})
	data.Items = items
	// fmt.Println(requestUrl)
	r := request(requestUrl, data)
	fmt.Println(r)
}
func request(url string, data interface{}) string {
	if _body, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		body := bytes.NewBuffer([]byte(_body))
		fmt.Println(body)
		res, err2 := http.Post(url, "application/json;charset=utf-8", body)
		if err2 != nil {
			fmt.Println(err2)
			os.Exit(1)
		}
		result, err3 := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(1)
		}
		// fmt.Printf("%s", result)
		return string(result)
	}
}

var _getTime = func() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func _getSignatureUrl(dataType, source string) string {
	newUrl := strings.Replace(_URL, "{type}", dataType, -1)
	newUrl = strings.Replace(newUrl, "{source}", source, -1)
	newUrl = newUrl + "&timestamp=" + _getTime()
	signature := _getHashHmacString(_TOKEN, newUrl)
	return newUrl + "&signature=" + signature
}
func _getHashHmacString(token, url string) string {
	mac := hmac.New(sha1.New, []byte(token))
	mac.Write([]byte(url))
	signature := fmt.Sprintf("%x", mac.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(signature))
}
