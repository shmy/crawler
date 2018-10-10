package entity

import (
	"encoding/xml"
)

/**
xml 解析
*/
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	List    List     `xml:"list"`
}
type List struct {
	XMLName xml.Name `xml:"list"`
	Videos  []*Video `xml:"video"`
}
type Video struct {
	XMLName  xml.Name `xml:"video"`
	Id       int      `xml:"id"`
	Name     string   `xml:"name"`
	Last     string   `xml:"last"`
	Note     string   `xml:"note"`
	Actor    string   `xml:"actor"`
	Director string   `xml:"director"`
	Des      string   `xml:"des"`
	Dl       Dl       `xml:"dl"`
}
type Dl struct {
	XMLName xml.Name `xml:"dl"`
	Dds     []*Dd    `xml:"dd"`
}

type Dd struct {
	XMLName xml.Name `xml:"dd"`
	Type    string   `xml:"flag,attr"`
	Value   string   `xml:",chardata"`
}
