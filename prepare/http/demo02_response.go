package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?theCityName=武汉"
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	fmt.Println("-----------------------------------------")
	fmt.Printf("response：%+v\n", resp)
	fmt.Println("-----------------------------------------")
	fmt.Printf("response.Body：%+v\n", resp.Body)
	fmt.Printf("response.Header：%+v\n", resp.Header)
	fmt.Printf("response.StatusCode：%+v\n", resp.StatusCode)
	fmt.Printf("response.Status：%+v\n", resp.Status)
	fmt.Printf("response.Request：%+v\n", resp.Request)
	fmt.Printf("response.Cookies：%+v\n", resp.Cookies())
}
