package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"
)

func TestRm(t *testing.T) {
	os.Chdir("ver_1.6.22.80")
	fn := "ver_1.6.22.80.zip"
	if _, err := os.Stat(fn); err == nil {
		os.Remove(fn)
	}
}

func TestFileName(t *testing.T) {
	dir := "ver_1.6.22.80"
	zipFile := "ver_1.6.22.80.zip"
	os.Chdir(dir)
	listFile := listDir("./", func(f fs.FileInfo) bool {
		if f.IsDir() {
			return false
		}
		if strings.HasSuffix(f.Name(), "cache.manifest") || f.Name() == zipFile {
			return false
		}
		return true
	})
	t.Log(len(listFile))
	//for _, fName := range listFile {
	//	t.Log(fName)
	//}

	compressZip(zipFile, listFile...)
	md5Str, sz, err := hashFileMd5AndSize(zipFile)
	if err != nil {
		t.Error(err)
	}
	t.Log(md5Str, sz)

}

func TestHashFileMd5(t *testing.T) {

	md5Str, sz, err := hashFileMd5AndSize("./ver_1.6.22.80.zip")
	if err != nil {
		t.Error(err)
	}
	t.Log(md5Str, sz)
}

//func TestLuaParse(t *testing.T) {
//	L := lua.NewState()
//	defer L.Close()
//	if err := L.DoFile("version.manifest"); err != nil {
//		t.Error(err)
//	}
//
//	lv := L.Get(-1)
//	if tbl, ok := lv.(*lua.LTable); ok {
//		// lv is LTable
//		//tbl.ForEach(tableCb
//		fmt.Println(tbl.RawGetString("updateCode").String())
//		//stageTbl := tbl.RawGetString("stage")
//	}
//}
//
//func tableCb(k, val lua.LValue) {
//	if tab, ok := val.(*lua.LTable); ok {
//		tab.ForEach(tableCb)
//	} else {
//		fmt.Println(k, val)
//	}
//}

func TestReadLine(t *testing.T) {
	//  (\w+)\s*=\s*"([\w|_|\.|\d|]+)"
	re := regexp.MustCompile(`(\w+)\s*=\s*"([\w|_|\.|\d|]+)"`)
	err := readFileLine("version.manifest", func(line string) {
		submatchArr := re.FindAllStringSubmatch(line, -1)
		if len(submatchArr) > 0 {
			t.Logf("%v", submatchArr)
		}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadVersionManifest(t *testing.T) {
	v, err := readVersionManifest("version.manifest")
	if err != nil {
		t.Fatal(err)
	}
	v.UpdateCode = "2021091002"
	luaTemplate := template.Must(template.New("").Parse(`
local m={
	updateCode = "{{.UpdateCode}}",
	stage={
        {{- range .Stage }}
{{"\t\t"}}{name="{{.Name}}",updateCode="{{.UpdateCode}}",code="{{.Code}}",size="{{.Size}}}"},
		{{- end }}
	}
}
return m
`))
	err = luaTemplate.Execute(os.Stdout, v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCode(t *testing.T) {
	code := "2021102501"
	sz := len(code)
	dateStr := string(code[:sz-2])
	vStr := string(code[sz-2:])
	i, err := strconv.Atoi(vStr)
	if err != nil {
		t.Error(err)
	}
	i++
	now := time.Now()
	dateStr = now.Format("20060102")
	newcode := fmt.Sprintf("%s%02d", dateStr, i)
	t.Logf("date:%s version:%s newVersion:%s", dateStr, vStr, newcode)
	var ssz int64 = 100000
	t.Log(strconv.FormatInt(ssz, 10))

}
