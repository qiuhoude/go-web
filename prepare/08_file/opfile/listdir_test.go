package opfile

import (
	"io/ioutil"
	"testing"
)

func TestListDir(t *testing.T) {
	listFile := listDir(".")
	for _, fName := range listFile {
		t.Log(fName)
	}
}

func listDir(dirName string) []string {
	var rt []string
	fileInfos, _ := ioutil.ReadDir(dirName)
	for _, fi := range fileInfos {
		fileName := dirName + "/" + fi.Name()
		rt = append(rt, fileName)
		if fi.IsDir() {
			rt = append(rt, listDir(fileName)...)
		}
	}
	return rt
}
