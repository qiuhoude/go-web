package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := fanIn(boring("hi"), boring("hello"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		//fmt.Println(rand.Intn(1e3))
	}
	fmt.Println("leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// 多个输入chan合并成一个
func fanIns(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(inputs))
		for _, i := range inputs {
			go func() {
				go func(i <-chan string) {
					for v := range i {
						out <- v
					}
					wg.Done()
				}(i)
			}()
		}
		wg.Wait()
		close(out)
	}()
	return out
}
