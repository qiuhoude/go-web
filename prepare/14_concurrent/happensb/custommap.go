package main

import "sync"

type Map struct {
	m map[int]int
	sync.Mutex
}

func (m *Map) Get(key int) (int, bool) {
	m.Lock()
	defer m.Unlock()
	i, ok := m.m[key]
	return i, ok
}

func (m *Map) Put(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}

func (m *Map) Len() int {
	return len(m.m)
}

/*
A: 不能编译
B: 可运行，无并发问题
C: 可运行，有并发问题  答案,需要将len也加上锁
D: panic
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := Map{m: make(map[int]int)}
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Put(i, i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Len()
		}
		wg.Done()
	}()
	wg.Wait()

}
