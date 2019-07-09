package main

import (
	"fmt"
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

/*
A: 不能编译
B: 可以编译，正确实现了单例
C: 可以编译，不正确的单例实现 答案,以为done==1会有线程读写冲突
D: 可以编译，但是程序运行会panic
*/
func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

/*
A: 不能编译
B: 可以编译，正确实现了单例
C: 可以编译，不正确的单例实现
D: 可以编译，但是程序运行会panic
*/

func main() {
	once := &Once{}
	for i := 0; i < 10; i++ {
		go once.Do(func() {
			fmt.Println("我执行了")
		})
	}
}
