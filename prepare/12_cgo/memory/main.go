package main

import (
	"sync"
	"time"
)

// 有bug 对count变量的访问并没有形成临界区
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 2)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
