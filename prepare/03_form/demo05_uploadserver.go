package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == http.MethodGet {
		curtime := time.Now().Unix() // unix 时间戳
		h := md5.New()
		curtimeStr := strconv.FormatInt(curtime, 10)
		fmt.Println(curtimeStr)
		io.WriteString(h, curtimeStr)
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token ->", token)
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println("token:", token)
		} else {
			//不存在token报错
			fmt.Println("token有误。。")
			//return
		}
		file, fileHeader, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", fileHeader.Header)
		f, err := os.Create("./test/" + fileHeader.Filename)
		//f, err := os.OpenFile("./test/"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 打开本地的文件
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file) //拷贝
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
