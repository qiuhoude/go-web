package main

import (
	"fmt"
	"sync"
	"time"
)

/*
提供相同服务的n个节点发送请求，只要任意一个服务节点返回结果，
我们就可以执行下面的业务逻辑，其它n-1的节点的请求可以被取消或者忽略
*/

//or函数可以处理n个channel,它为每个channel启动一个goroutine，
// 只要任意一个goroutine从channel读取到数据，输出的channel就被关闭掉了。
func or(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range chans {
			go func(ch <-chan interface{}) {
				select {
				case <-ch:
					once.Do(func() { // 为了避免并发关闭输出channel的问题，关闭操作只执行一次
						close(out)
					})
				case <-out:
				}
			}(c)
		}
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(10*time.Second),
		sig(5*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(10*time.Second),
		sig(01*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
