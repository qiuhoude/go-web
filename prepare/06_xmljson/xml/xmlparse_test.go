package xml

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestParseServerXml(t *testing.T) {
	file, err := os.Open("servers.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestParseStudentXml(t *testing.T) {
	file, err := os.Open("students.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	v := Recurlystudents{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
