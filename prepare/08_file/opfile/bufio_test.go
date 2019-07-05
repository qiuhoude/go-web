package opfile

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBufioDemo(t *testing.T) {
	bufioDemo()
}
func bufioDemo() {

	/*
		   bufio:高效io读写
			   buffer缓存
			   io：input/output

		   将io包下的Reader，Write对象进行包装，带缓存的包装，提高读写的效率

			   ReadBytes()
			   ReadString()
			   ReadLine()

	*/

	fileName := "aa.txt"
	file, _ := os.Open(fileName) //看作是io包下的Reader，Write的实现
	defer file.Close()
	r1 := bufio.NewReader(file) //构建带缓存的Reader对象：bufio.Reader
	fmt.Printf("%T\n", r1)
	data, flag, err := r1.ReadLine()
	fmt.Println(err)
	fmt.Println(flag)
	fmt.Println(data)
	fmt.Println(string(data))

	//s1, _ := r1.ReadString('\n')
	//fmt.Println(s1)
	//s1, _ = r1.ReadString('\n')
	//fmt.Println(s1)
	//s1, _ = r1.ReadString('\n')
	//fmt.Println(s1)
	//s1, _ = r1.ReadString('\n')
	//fmt.Println(s1)
	//
	//for {
	//	s1, err := r1.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("读取完毕。。")
	//		break
	//	}
	//	fmt.Print(s1)
	//}

	data, _ = r1.ReadBytes('\n')
	fmt.Println(string(data))

	s2 := ""
	fmt.Scanln(&s2)
	fmt.Println(s2)

	r2 := bufio.NewReader(os.Stdin)
	s2, _ = r2.ReadString('\n')
	fmt.Println(s2)

}
