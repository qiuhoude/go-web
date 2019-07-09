package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}

/*
A: 不能编译
B: 输出 1 答案
C: 输出 0
D: panic
*/
func main() {
	go f()
	c <- 0
	print(a)
}
