package main

import (
	"fmt"
	"sync"
	"time"
)

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}

/*
A. 不能编译；
B. 输出 1；
C. 输出 0；
D. panic；

B。能正确输出，不过主协程会阻塞 f() 函数的执行。
*/
func demo1() {
	go f()
	c <- 0
	fmt.Println(a)
}

/*
A. 不能编译；
B. 输出 1, 1；
C. 输出 1, 2；
D. fatal error；

D。加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。

*/
type MyMutex struct {
	count int
	sync.Mutex
}

func demo2() {
	//fmt.Errorf()
	var mu MyMutex

	mu.Lock()
	var mu1 = mu
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

/*
A. 不能编译；
B. 无输出，正常退出；
C. 程序 hang 住；
D. panic；

D。WaitGroup 在调用 Wait() 之后不能再调用 Add() 方法的。
*/
func demo3() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}

func main() {
	//demo1()
	//demo2()
	//demo3()

}
