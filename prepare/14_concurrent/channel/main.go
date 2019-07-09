package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
A: 不能编译
B: 一段时间后总是输出 #goroutines: 1
C: 一段时间后总是输出 #goroutines: 2 答案
D: panic
*/
func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
		fmt.Println("结束1")
	}()

	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
		fmt.Println("结束2")
	}(ch)

	c := time.Tick(1 * time.Second)

	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
