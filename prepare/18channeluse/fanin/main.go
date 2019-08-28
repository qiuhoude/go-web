package main

import (
	"reflect"
	"sync"
)

/*
扇入模式(FanIn)是将多个同样类型的输入channel合并成一个同样类型的输出channel，也就是channel的合并。

*/

// https://github.com/campoy/justforfunc/blob/master/27-merging-chans/main.go
//1.Goroutine方式  每个channel起一个goroutine。
func fanIn(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan interface{}) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()
	return out
}

//2. Reflect 方式
// 利用反射库针对select语句的处理合并输入channel。在输入channel读取比较均匀的时候比较有效，否则性能比较低下
func fanInReflect(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
			for len(cases) > 0 {
				i, v, ok := reflect.Select(cases)
				if !ok { //remove this case
					cases = append(cases[:i], cases[i+1:]...)
					continue
				}
				out <- v.Interface()
			}
		}
	}()
	return out
}

//3.递归方式
//这种方式虽然理解起来不直观，但是性能还是不错的(输入channel不是很多的情况下递归层级不会很高，不会成为瓶颈)
func fanInRec(chans ...<-chan interface{}) <-chan interface{} {
	switch len(chans) {
	case 0:
		return nil
	case 1:
		return chans[0]
	case 2:
		return mergeTwo(chans[0], chans[1])
	default:
		m := len(chans) / 2
		return mergeTwo(
			fanInRec(chans[:m]...), // 左边
			fanInRec(chans[m:]...)) // 右边
	}
}
func mergeTwo(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
