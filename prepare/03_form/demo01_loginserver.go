package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello my route!") //这个写入到w的是输出到客户端的

}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == http.MethodPost {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	} else if r.Method == http.MethodGet {
		t, err := template.ParseFiles(`login.html`) // 解析文件模板
		if err != nil {
			log.Println(err)
			return
		}
		t.Execute(w, "登陆页面")
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}
