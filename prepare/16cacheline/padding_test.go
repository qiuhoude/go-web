package cacheline

import (
	"sync/atomic"
	"testing"
)

//https://colobu.com/2019/01/24/cacheline-affects-performance-in-go/

/*
可以使用intel-go/cpuid获取CPU的cacheline的大小
https://github.com/intel-go/cpuid

也提供了一个CacheLinePad struct用来padding 	cpu.CacheLinePad{}
https://github.com/golang/sys/blob/master/cpu/cpu.go
*/

type NoPad struct {
	//_ cpu.CacheLinePad // 缓存行
	a uint64
	b uint64
	c uint64
}

func (np *NoPad) Increase() {
	atomic.AddUint64(&np.a, 1)
	atomic.AddUint64(&np.b, 1)
	atomic.AddUint64(&np.c, 1)
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	b   uint64
	_p2 [8]uint64
	c   uint64
	_p3 [8]uint64
}

func (p *Pad) Increase() {
	atomic.AddUint64(&p.a, 1)
	atomic.AddUint64(&p.b, 1)
	atomic.AddUint64(&p.c, 1)
}

func BenchmarkPad_Increase(b *testing.B) {
	pad := &Pad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pad.Increase()
		}
	})
}

func BenchmarkNoPad_Increase(b *testing.B) {
	nopad := &NoPad{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nopad.Increase()
		}
	})
}
