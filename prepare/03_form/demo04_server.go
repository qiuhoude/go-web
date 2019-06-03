package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

//防止多次递交表单

func login4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法

	if r.Method == http.MethodGet {
		curtime := time.Now().Unix() // unix 时间戳
		h := md5.New()
		curtimeStr := strconv.FormatInt(curtime, 10)
		fmt.Println(curtimeStr)
		io.WriteString(h, curtimeStr)
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token ->", token)
		t, _ := template.ParseFiles("test.gtpl")
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println("token:", token)
		} else {
			//不存在token报错
			fmt.Println("token有误。。")
		}
		fmt.Println("username length:", len(r.Form.Get("username")))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}

}

func main() {
	http.HandleFunc("/login4", login4)       //设置访问的路由
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
