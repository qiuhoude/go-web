package unionfind

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkUnionFind(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	size := b.N
	uf := NewUnionFind(size)
	for i := 0; i < b.N; i++ {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.UnionElements(a, b)
	}
	for i := 0; i < b.N; i++ {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.IsConnected(a, b)
	}
}

func BenchmarkUnionFindSize(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	size := b.N
	uf := NewUnionFindOpSize(size)
	for i := 0; i < b.N; i++ {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.UnionElements(a, b)
	}
	for i := 0; i < b.N; i++ {
		a := r.Intn(size)
		b := r.Intn(size)
		uf.IsConnected(a, b)
	}
}
