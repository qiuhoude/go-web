package opfile

import (
	"io/ioutil"
	"testing"
)

func TestListDir(t *testing.T) {
	listDir("/project")
}

func listDir(dirName string) {
	fileInfos, _ := ioutil.ReadDir(dirName)
	for _, fi := range fileInfos {
		fileName := dirName + "/" + fi.Name()
		//fmt.Println(fileName)
		if fi.IsDir() {
			listDir(fileName)
		}
	}
}
