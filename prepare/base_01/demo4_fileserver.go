package main

import (
	"log"
	"net/http"
)

func main() {
	/*
			FileServer：
	
			1.www.xx.com/ 根路径 直接使用
		　　http.Handle("/", http.FileServer(http.Dir("/tmp")))
		　　2.www.xx.com/c/ 带有请求路径的 需要添加函数
		　　http.Handle("/c/", http.StripPrefix("/c/", http.FileServer(http.Dir("/tmp"))))
	*/
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/f/", http.StripPrefix("/f/", http.FileServer(http.Dir("/log"))))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
