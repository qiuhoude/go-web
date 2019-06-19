package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("unsafe_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name()) //aa.txt
	//是否是目录
	fmt.Println(fileInfo.IsDir()) //false
	//大小
	fmt.Println(fileInfo.Size()) //1
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-
	//修改时间
	fmt.Println(fileInfo.ModTime())
}
