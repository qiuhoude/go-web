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
		"username": {"houdeqiu"},
		"password": {"123456"},
	}

	requestBody := bytes.NewBufferString(param.Encode())
	resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", requestBody)

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
}
