package main

import "fmt"

/*
这种模式是我们经常使用的一种模式， 通过一个信号channel(done)来控制(取消)输入channel的处理。
一旦从done channel中读取到一个信号，或者done channel被关闭， 输入channel的处理则被取消。
这个模式提供一个简便的方法，把done channel 和 输入 channel 融合成一个输出channel。

就会context的前身,通过done进行取消整个运行中的任务
*/

func orDone(done <-chan struct{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func main() {
	ch := make(chan interface{})
	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for v := range orDone(nil, ch) {
		fmt.Printf("%v ", v)
	}
}
