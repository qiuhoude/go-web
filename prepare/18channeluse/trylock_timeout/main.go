package main

import (
	"context"
	"fmt"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}
func (m *Mutex) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}

// context的方式
func (m *Mutex) TryLock2(timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	select {
	case <-m.ch:
		cancel()
		return true
	case <-ctx.Done():
		fmt.Println("err -> ", ctx.Err())
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}
func main() {
	m := NewMutex()
	ok := m.TryLock(time.Second)
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock(time.Second)
	fmt.Printf("locked %v\n", ok)
}
