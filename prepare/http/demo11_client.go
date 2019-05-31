package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	urlStr := "http://localhost:8000/login"
	param := url.Values{
		"username": {"qiuenqi"},
		"password": {"123456"},
	}
	requestBody := bytes.NewBufferString(param.Encode())
	client := new(http.Client)
	request, err := http.NewRequest(http.MethodPost, urlStr, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	/*
		cookie的添加
		方式一：
			request.Header.Set("Cookie", "name=hanru")

		方式二：
			request.AddCookie(Cookie)

	*/
	cookId := &http.Cookie{Name: "userId", Value: "1234"}
	cookName := &http.Cookie{Name: "name", Value: "kongyixueyuan"}
	request.AddCookie(cookId)
	request.AddCookie(cookName)
	//使用text/html就会出错
	//application/x-www-form-urlencoded
	request.Header.Set("Content-Type", "text/html")
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("resp:%+v\n", resp)
	fmt.Printf("resp.Header:%+v\n", resp.Header)
	fmt.Printf("resp.Cookies:%+v\n", resp.Cookies())
	fmt.Printf("request.Header:%+v\n", resp.Header)
	fmt.Printf("request.Cookies:%+v\n", resp.Cookies())
}
