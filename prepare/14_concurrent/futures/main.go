package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://www.baidu.com"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes) err

}
