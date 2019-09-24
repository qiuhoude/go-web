package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
扇出模式(FanOut)是将一个输入channel扇出为多个channel。

扇出行为至少可以分为两种：

1. 从输入channel中读取一个数据，发送给每个输入channel，这种模式称之为Tee模式; 个人理解:有点类似于redis中的订阅发布
2. 从输入channel中读取一个数据，在输出channel中选择一个channel发送
*/

//1.
// 一个输入 ,多个输出
func fanOut1(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()
		for v := range ch {
			v := v // 重新声明,避免闭包应用问题
			for i := 0; i < len(out); i++ {
				i := i
				if async {
					go func() {
						out[i] <- v
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()
}

// 递归方式
func fanOut1Reflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()
		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend // 发送选择
		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Chan = reflect.ValueOf(out[i])
				cases[i].Send = reflect.ValueOf(v)
			}
			// 选择
			for _ = range cases {
				i, _, _ := reflect.Select(cases)
				cases[i].Chan = reflect.ValueOf(nil)
			}
		}

	}()
}

// 2.
// 一个输入, 选其中一个输出 分布模式
func fanOut2(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		// roundrobin
		var i = 0
		var n = len(out)
		for v := range ch {
			v := v
			out[i] <- v
			i = (i + 1) % n
		}
	}()
}

func fanOut2Reflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])

		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			_, _, _ = reflect.Select(cases)
		}
	}()
}

func asStream(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)

		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}

	}()
	return s
}

func runFanout1() {
	source := asStream(nil)
	channels := make([]chan interface{}, 5)

	fmt.Println("fanOut1")
	for i := 0; i < 5; i++ {
		channels[i] = make(chan interface{})
	}
	fanOut1(source, channels, false)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("channel#%d: %v\n", j, <-channels[j])
		}
	}

	fmt.Println("\nfanOut1 By Reflect")
	source = asStream(nil)
	for i := 0; i < 5; i++ {
		channels[i] = make(chan interface{})
	}
	fanOut1Reflect(source, channels)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("channel#%d: %v\n", j, <-channels[j])
		}
	}
}
func runFanout2() {
	done := make(chan struct{})
	source := asStream(done)
	channels := make([]chan interface{}, 5)

	fmt.Println("fanOut")
	for i := 0; i < 5; i++ {
		channels[i] = make(chan interface{})
	}
	fanOut2(source, channels)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			for j := 0; j < 5; j++ {
				v, ok := <-channels[i]
				if ok {
					fmt.Printf("channel#%d: %v\n", i, v)
				}

			}
		}()
	}
	time.Sleep(time.Second)
	close(done)

	fmt.Println("fanOut By Reflect")
	done = make(chan struct{})
	source = asStream(done)
	for i := 0; i < 5; i++ {
		channels[i] = make(chan interface{})
	}
	fanOut2Reflect(source, channels)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			for j := 0; j < 5; j++ {
				v, ok := <-channels[i]
				if ok {
					fmt.Printf("channel#%d: %v\n", i, v)
				}
			}
		}()
	}
	time.Sleep(time.Second)
	close(done)
}

func main() {

	//runFanout1()
	runFanout2()
}
