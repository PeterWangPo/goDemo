package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type jsonTime time.Time

//实现它的json序列化方法
func (this jsonTime) MarshalJSON() ([]byte, error) {
    stamp := fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
    return []byte(stamp), nil
}

type Test struct {
    Date  jsonTime `json:"date"`
    Name  string   `json:"name"`
    State bool     `json:"state"`
}

func main() {
    var t = Test{}
    t.Date = jsonTime(time.Now())
    t.Name = "Hello World"
    t.State = true
    body, _ := json.Marshal(t)
    fmt.Println(string(body))
}
