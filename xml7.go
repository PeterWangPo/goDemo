package main

import (
	"encoding/xml"

	"log"
)

func main() {
	type CDATA struct {
		Text string `xml:",cdata"`
	}

	type TextMsg struct {
		XMLName    xml.Name `xml:"xml"`
		ToUserName CDATA    `xml:"to_User"`
	}

	msg := TextMsg{
		ToUserName: CDATA{"userId"},
	}

	xmlStr, err := xml.Marshal(msg)
	if err != nil {
		log.Printf("generate xml err:%s, xml param:%+v", err, xmlStr)
	}
	log.Printf("xml:%s, ", xmlStr)

}
