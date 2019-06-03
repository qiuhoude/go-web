package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/**
	方式二 调用client的API

	1.先生成client，
	2.之后用client.get/post..

	还额外添加了req.Header.Set("Content-Type", bodyType)...
	*/
	client := http.Client{}
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=北京"
	resp, err := client.Get(urlStr)
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
	fmt.Printf("%+v", resp)
}
