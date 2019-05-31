package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	http.ListenAndServe(":8000", mux)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析表单
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("body : ", string(body))
	fmt.Printf("r : %+v\n", r)
	fmt.Printf("Header : %+v\n", r.Header)
	fmt.Printf("Content-Type : %+v\n", r.Header["Content-Type"])
	fmt.Printf("Cookies : %+v\n", r.Cookies())

	if len(r.Form["username"]) > 0 {
		username := r.Form["username"][0]
		fmt.Println("username:", username)
	}
	if len(r.Form["password"]) > 0 {
		password := r.Form["password"][0]
		fmt.Println("password:", password)
	}

	io.WriteString(w, "登录成功")
}
