package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	url := "http://localhost:8000/upload"
	filename := "./1.jpg"
	postFile(url, filename)
}

func postFile(url, filename string) error {
	bodybuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodybuf)
	bodyWriter.WriteField("token", "03126ee75b098a46f395694b20ac2eqiu")
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 打开文件句柄
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 拷贝
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// post请求
	resp, err := http.Post(url, contentType, bodybuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
