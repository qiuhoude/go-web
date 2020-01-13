package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func commonPoolDemo(runTimes int) {
	defer ants.Release()

	var wg sync.WaitGroup

	// Use the common pool.
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum) // 提交任务
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}

func funcPoolDemo(runTimes int) {
	var wg sync.WaitGroup
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	}, ants.WithPreAlloc(true))
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

func customPoolDemo(runTimes int) {
	var wg sync.WaitGroup
	p, _ := ants.NewPool(1, ants.WithPreAlloc(true))
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		param := int32(i)
		_ = p.Submit(func() {
			myFunc(param)
			//demoFunc()
			wg.Done()
		}) // 提交任务
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

func main() {
	//commonPoolDemo(100)
	//funcPoolDemo(100)
	customPoolDemo(100)
}
