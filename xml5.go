package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Config map[string]map[string]string

var config = make(Config)

type Class struct {
	XMLName xml.Name  `xml:"class"`     //xml元素名称
	Id      int       `xml:"id,attr"`   //
	Pers    []Persion `xml:"persion"`   //
	Desc    string    `xml:",innerxml"` //
}

type Persion struct {
	XMLName xml.Name `xml:"persion"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func (self *Config) LoadXmlFile(filename, node string) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	return self.Load(fd, node)
}
func (self *Config) Load(r io.Reader, node string) error {
	mynode := false
	if _, ok := (*self)[node]; !ok {
		(*self)[node] = make(map[string]string)
	}
	decoder := xml.NewDecoder(r)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		switch value := token.(type) {
		case xml.StartElement:
			switch {
			case value.Name.Local == node:
				mynode = true
			case mynode == true:
				tb, err := decoder.Token()
				if err != nil {
					continue
				}
				switch tv := tb.(type) {
				case xml.CharData:
					(*self)[node][value.Name.Local] = string(tv)
				}
			}
		case xml.EndElement:
			if value.Name.Local == node {
				mynode = false
			}
		}
	}
	return nil
}

func main() {
	var ps = make([]Persion, 0)
	for i := 0; i < 5; i++ {
		var tp Persion
		tp.Id = i
		tp.Name = "persion" + strconv.Itoa(i)
		tp.Age = 20 + i
		ps = append(ps, tp)
	}
	var c = Class{Id: 1, Pers: ps}
	var res, _ = xml.MarshalIndent(c, "", "    ")
	fmt.Println(c)
	fmt.Println(string(res))
	config.LoadXmlFile("config.xml", "global")
	fmt.Println(config)
}
