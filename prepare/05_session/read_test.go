package main

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	content, err := ioutil.ReadFile("login.html")
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("%s \n", content)
	}
}
