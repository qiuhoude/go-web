package main

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	gameDirStr     string
	inManifestStr  string
	outManifestStr string
)

func main() {
	flag.StringVar(&gameDirStr, "gameDir", "", "游戏目录")
	flag.StringVar(&inManifestStr, "inManifest", "version.manifest", "版本配置文件")
	flag.StringVar(&outManifestStr, "outManifest", "version.manifest.new", "生成新版本配置文件")
	flag.Parse()

	// 压缩文件
	err := os.Chdir(gameDirStr) // cd gameDir/
	if err != nil {
		log.Fatal(err)
	}
	zipFile := gameDirStr + ".zip"
	listFile := listDir("./", func(f fs.FileInfo) bool {
		if f.IsDir() ||
			strings.HasSuffix(f.Name(), "cache.manifest") ||
			f.Name() == zipFile {
			return false
		}
		return true
	})

	compressZip(zipFile, listFile...)

	// md5 和 size
	md5Str, sz, err := hashFileMd5AndSize(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chdir("../")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("计算md5值完成 %s -> md5:%s size:%v", zipFile, md5Str, sz)

	// 读取 VersionManifest
	vm, err := readVersionManifest(inManifestStr)
	if err != nil {
		log.Fatal(err)
	}

	// 新的stage数据
	newStage := Stage{
		Name: zipFile,
		Code: strings.ToUpper(md5Str),
		Size: strconv.FormatInt(sz, 10),
	}

	// 计算每日版本号
	now := time.Now()
	curDateStr := now.Format("20060102")
	todayVersion := 0
	if len(vm.Stage) > 0 {
		lastStage := vm.Stage[len(vm.Stage)-1]
		n := len(lastStage.UpdateCode) - 1
		dateStr := string(lastStage.UpdateCode[:n-2])
		vStr := string(lastStage.UpdateCode[n-2:])
		if curDateStr == dateStr {
			i, err := strconv.Atoi(vStr)
			if err != nil {
				log.Fatal(err)
			}
			todayVersion = i
		}
	}
	todayVersion++
	newStage.UpdateCode = fmt.Sprintf("%s%02d", curDateStr, todayVersion)
	vm.UpdateCode = newStage.UpdateCode
	vm.Stage = append(vm.Stage, newStage) // 添加到数据中
	err = outPutVersionManifestFile(vm, outManifestStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("成功完成...")

}

type fileFilter func(fileName fs.FileInfo) bool

type VersionManifest struct {
	UpdateCode string
	Stage      []Stage
}

type Stage struct {
	Name       string
	UpdateCode string
	Code       string
	Size       string
}

func listDir(dirName string, filter fileFilter) []string {
	var rt []string
	err := filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filter != nil && filter(info) {
			rt = append(rt, path)
		}
		return nil

	})
	if err != nil {
		return nil
	}

	return rt
}
func rmFile(zipFileName string) {
	if _, err := os.Stat(zipFileName); err == nil {
		_ = os.Remove(zipFileName)
		log.Printf("删除文件 %v", zipFileName)
	}
}

// 压缩文件
func compressZip(zipFileName string, files ...string) {
	//rmFile(zipFileName)
	file, err := os.OpenFile(zipFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//file, err := os.Create(zipFileName)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	zipw := zip.NewWriter(file)
	defer zipw.Close()
	totalSize := len(files)
	for i, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
		fmt.Printf("\r>>>progress 正在压缩文件 (%d/%d) ", i+1, totalSize)
	}
	fmt.Println()
	fmt.Printf("压缩文件完成 %s \n", zipFileName)

}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}
	return nil
}

func hashFileMd5AndSize(filePath string) (md5Str string, filesz int64, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	info, err := file.Stat()
	if err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	filesz = info.Size()
	return
}

func readFileLine(path string, doFunc func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		doFunc(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func readVersionManifest(path string) (v VersionManifest, err error) {
	//  (\w+)\s*=\s*"([\w|_|\.|\d|]+)"
	//codeRe := regexp.MustCompile(`updateCode = "([\w|_|\.|\d|]+)"`) // 外层的updateCode
	stageRe := regexp.MustCompile(`(\w+)="([\w|_|\.|\d|]+)"`)
	v = VersionManifest{}
	err = readFileLine(path, func(line string) {
		arr := stageRe.FindAllStringSubmatch(line, -1)
		if len(arr) > 0 {
			var s = Stage{}
			for _, d := range arr {
				switch d[1] {
				case "name":
					s.Name = d[2]
				case "updateCode":
					s.UpdateCode = d[2]
				case "code":
					s.Code = d[2]
				case "size":
					s.Size = d[2]
				}
			}
			v.Stage = append(v.Stage, s)
		}
	})
	if err != nil {
		return
	}
	return
}

func outPutVersionManifestFile(v VersionManifest, fp string) error {
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	luaTemplate := template.Must(template.New("").Parse(`local m={
	updateCode = "{{.UpdateCode}}",
	stage={
        {{- range .Stage }}
{{"\t\t"}}{name="{{.Name}}",updateCode="{{.UpdateCode}}",code="{{.Code}}",size="{{.Size}}"},
		{{- end }}
	}
}
return m`))
	err = luaTemplate.Execute(file, v)
	return err
}
