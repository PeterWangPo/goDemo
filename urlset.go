package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

type urlset struct {
	Urlset url `json:"urlset"`
}
type url struct {
	Url []urlType `json:"url"`
}

type urlType struct {
	Loc  string   `json:"loc"`
	Data dataType `json:"data"`
}
type dataType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	// var urlsetdata urlset
	jsonstr := `{"urlset":{"url":[{"loc":"1","data":{"id":1,"name":"1"}},{"loc":"2","data":{"id":2,"name":"2"}}]}}`
	// detail = append(detail, urls)
	// urlsetdata.Urlset = urlDatas
	// var box map[string]urlset
	body, err := json.Marshal(jsonstr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	js, err := simplejson.NewJson(body)
	fmt.Println(js)
}
