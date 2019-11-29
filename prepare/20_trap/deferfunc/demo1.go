package main

import "fmt"

func deferFunReturn() (result int) {
	i := 0
	defer func() {
		//i++ // 不会改变返回值
		result++ // 会改变返回值  延迟函数可能操作主函数的具名返回
	}()
	return i
}

func deferFunReturn2() int {
	i := 0
	defer func() {
		i++ // 不会改变返回值
	}()
	return i
}

func main() {
	fmt.Println(deferFunReturn())
	fmt.Println(deferFunReturn2())
	go Dived(0)

	//Dived2(0) // 会 panic

}

func NoPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
	}
}

func Dived2(n int) {
	/*
	   以下三个条件会让recover()返回nil:
	   1. panic时指定的参数为nil；（一般panic语句如panic("xxx failed...")）
	   2. 当前协程没有发生panic；
	   3. recover没有被defer方法直接调用
	*/
	defer func() {
		NoPanic() // 将会失败
	}()

	fmt.Println(1 / n)
}

func Dived(n int) {
	defer NoPanic()

	fmt.Println(1 / n)
}
