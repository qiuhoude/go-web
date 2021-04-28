package main

import (
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestHello(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	t.Log(item.V)
}
