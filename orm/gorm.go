package main

import (
    "github.com/jinzhu/gorm"
    "time"
    "fmt"
    "strings"
    "math/rand"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

type tmpDevice struct {
    ID int64
    DeviceType string
    OriginalDeviceId string
    Md5DeviceId string
    LabelTag int32
    //LyTime time.Time
    LyTime int64
}

func (t tmpDevice) TableName() string {
    return "tmp_device"
}

func main() {
    //create
    connect()
    t := tmpDevice{
        DeviceType:"IDFA_MD5_32",
        OriginalDeviceId:strings.Repeat("a",rand.New(rand.NewSource(time.Now().Unix())).Int() % 32),
        Md5DeviceId:strings.Repeat("a",32),
        LabelTag:0,
        LyTime:time.Now().Unix(),
    }
    fmt.Println("t:", t)
    //r := conn.NewRecord(t)
    r := conn.Create(&t)
    fmt.Println("r:",r)
}

func connect() {
    dns := buildMysqlDNS()
    fmt.Println("dns:", dns)
    handle, err := gorm.Open("mysql", dns)
    //defer handle.Close()
    if err != nil {
        panic(err)
    }
    conn = handle
}

func buildMysqlDNS() string {
    return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		"xxx",
		"xxxx",
		"127.0.0.1",
		"3306",
		"xxx",
		"utf8",
		"True",
		"Local",
	)
}
