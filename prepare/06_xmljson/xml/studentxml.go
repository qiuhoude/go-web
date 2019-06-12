package xml

import "encoding/xml"

type Recurlystudents struct {
	XMLName     xml.Name  `xml:"students"`
	Version     string    `xml:"version,attr"`
	Students    []student `xml:"student"`
	Description string    `xml:",innerxml"`
}
type student struct {
	XMLName     xml.Name      `xml:"student"`
	StudentName string        `xml:"studentName"`
	Age         int           `xml:"age"`
	Sex         string        `xml:"sex"`
	Books       Recurlybookss `xml:"books"`
}

type Recurlybookss struct {
	XMLName     xml.Name `xml:"books"`
	Version     string   `xml:"version,attr"`
	Books       []book   `xml:"book"`
	Description string   `xml:",innerxml"`
}

type book struct {
	XMLName  xml.Name `xml:"book"`
	BookName string   `xml:"bookName"`
	Price    string   `xml:"price"`
}
