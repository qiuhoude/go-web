package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	// 防止 XSS Cross Site Scripting  跨站脚本攻击 使用 template.HTMLEscapeString 输出
	fmt.Println("username:", template.HTMLEscapeString(username)) //输出到服务器端
	fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
	template.HTMLEscape(w, []byte(username)) //输出到客户端
}

func login3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	fmt.Printf(username)
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(w, "T", username)
	err = t.ExecuteTemplate(w, "T", template.HTML(username))

	//如果转义失败 抛出对应错误 终止程序
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	http.HandleFunc("/login2", login2)
	http.HandleFunc("/login3", login3)
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
