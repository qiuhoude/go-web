package opfile

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	userFile := "aa.txt"
	//fout, err := os.Create(userFile)
	fout, err := os.OpenFile(userFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}

func TestReadFile(t *testing.T) {
	/*
		读取文件：
		1.打开文件
		2.读取文件
			file.Read([]byte)-->n,err
			从文件中开始读取数据，存入到byte切片中，返回值n是本次实际读取的数据量
				如果读取到文件末尾，n为0,err为EOF：end of file
		3.关闭文件
	*/
	userFile := "aa.txt"
	file, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到文件末尾了，结束读取操作。。")
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
