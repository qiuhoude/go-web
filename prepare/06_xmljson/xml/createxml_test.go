package xml

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestCreateXml(t *testing.T) {
	v := Recurlyservers{Version: "2"}
	v.Svs = append(v.Svs, server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	out, err := xml.MarshalIndent(v, " ", "	")
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(out)
}

type human struct {
	XMLName xml.Name `xml:"human"`
	Name    string   `xml:"info>name"`
	Age     string   `xml:"info>age"`
}

func TestCreateXml2(t *testing.T) {
	v := human{Name: "houde", Age: "18"}
	out, err := xml.MarshalIndent(v, " ", "	")
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(out)
}
