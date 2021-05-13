package main

import (
	"context"
	"errors"
	"github.com/reactivex/rxgo/v2"
	"strconv"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	if item.Error() {
		t.Error(item.E)
	}
	t.Log(item.V)
}

func TestForeach(t *testing.T) {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4)()

	<-observable.
		Map(func(context context.Context, i2 interface{}) (i interface{}, e error) {
			i = strconv.Itoa(i2.(int)) + "___"
			return
		}, rxgo.WithCPUPool()).
		ForEach(func(v interface{}) {
			t.Logf("received: %v\n", v)
		}, func(err error) {
			t.Logf("error: %e\n", err)
		}, func() {
			t.Logf("observable is closed")
		})
}

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func TestRealWorld(t *testing.T) {
	/*// Create the input channel
	ch := make(chan rxgo.Item)
	// Data producer
	go producer(ch)

	// Create an Observable
	observable := rxgo.FromChannel(ch)*/
}

func TestHotObservables(t *testing.T) {
	ch := make(chan rxgo.Item)

	go func() {
		for i := 0; i < 3; i++ {
			//time.Sleep(1 * time.Second)
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)

	// First Observer
	for item := range observable.Observe() {
		t.Log(item.V, "_1")
	}

	// Second Observer
	for item := range observable.Observe() {
		t.Log(item.V, "_2")
	}
	t.Log("end")
}

func TestColdObservables(t *testing.T) {

	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			//time.Sleep(1 * time.Second)
			ch <- rxgo.Of(i)
		}
	}, func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 3; i < 6; i++ {
			//time.Sleep(1 * time.Second)
			ch <- rxgo.Of(i)
		}
	}})
	// First Observer
	for item := range observable.Observe() {
		t.Log(item.V, "_1")
	}
	//time.Sleep(1 * time.Second)

	// Second Observer
	for item := range observable.Observe() {
		t.Log(item.V, "_2")
	}
	t.Log("end")
}

func TestConnectable(t *testing.T) {
	ch := make(chan rxgo.Item)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy()) // 发布模式,下面两个observable都可以收到

	observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	}).DoOnNext(func(i interface{}) {
		t.Logf("First observer: %d\n", i)
	})

	observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return i.(int) * 2, nil
	}).DoOnNext(func(i interface{}) {
		t.Logf("Second observer: %d\n", i)
	})
	// 只有 Connect observable才能发生数据
	_, _ = observable.Connect(context.Background())
	//disposable()

	for c := range observable.Observe() {
		t.Logf("%T %v", c, c)
	}

}

func TestSupplier(t *testing.T) {
	observable := rxgo.Start([]rxgo.Supplier{func(_ context.Context) rxgo.Item {
		return rxgo.Of(1)
	}, func(_ context.Context) rxgo.Item {
		return rxgo.Of(2)
	}})
	for item := range observable.Observe() {
		t.Logf("%v", item.V)
	}
}

func TestBuffer(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4)()

	observable = observable.BufferWithCount(3)

	for item := range observable.Observe() {
		t.Logf("%v", item.V)
	}
}
func TestJoin(t *testing.T) {
	observable := rxgo.Just(
		map[string]int64{"tt": 1, "V": 1},
		map[string]int64{"tt": 4, "V": 2},
		map[string]int64{"tt": 7, "V": 3},
	)()
	observableRight := rxgo.Just(
		map[string]int64{"tt": 2, "V": 5},
		map[string]int64{"tt": 3, "V": 6},
		map[string]int64{"tt": 5, "V": 7},
	)()

	observable.Join(func(ctx context.Context, l interface{}, r interface{}) (interface{}, error) {
		return map[string]interface{}{
			"l": l,
			"r": r,
		}, nil
	}, observableRight, func(i interface{}) time.Time {
		return time.Unix(i.(map[string]int64)["tt"], 0)
	}, rxgo.WithDuration(1))

	for item := range observable.Observe() {
		t.Logf("%v", item.V)
	}
	time.Sleep(5 * time.Second)
}
