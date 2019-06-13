package opfile

import (
	"fmt"
	"io"
	"os"
)

// 返回拷贝数量
func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024, 1024) //拷贝数据
	n := -1                        //读取的数据量
	total := 0
	for {
		if n, err = file1.Read(bs); err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。。")
			break
		} else if err != nil {
			fmt.Println("报错了。。。")
			return total, err
		} else {
			total += n
			file2.Write(bs[:n])
		}
	}
	return total, nil
}

func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}
