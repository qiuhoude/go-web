package minheap

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Task struct {
	Name     string
	Priority int
}

func (t *Task) CompareTo(e Comparable) int {
	o, _ := e.(*Task)
	return t.Priority - o.Priority
}

func generatorTask() []Comparable {
	rand.Seed(time.Now().Unix())
	const size = 10
	d := make([]Comparable, 0, size)
	for i := 0; i < size; i++ {
		p := rand.Intn(300)
		d = append(d, &Task{
			Name:     "task-" + strconv.Itoa(i),
			Priority: p,
		})
	}
	return d
}

func TestHeap_Remove(t *testing.T) {
	tasks := generatorTask()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
	fmt.Println("----------------")
	h := NewHeap(generatorTask())
	remove := h.Remove(tasks[1], func(e, b Comparable) bool {
		t1, ok1 := e.(*Task)
		t2, ok2 := e.(*Task)
		if ok1 && ok2 && t1 == t2 {
			fmt.Println("删除了 ", t1.Name)
			return true
		}
		return false
	})

	fmt.Println("删除 ", remove)
	for h.Len() != 0 {
		e := h.Poll()
		if t, ok := e.(*Task); ok {
			fmt.Println(t)
		}
	}
}

func TestHeap_Add(t *testing.T) {
	h := NewHeap(generatorTask())
	cnt := 0
	for h.Len() != 0 {
		cnt++
		pe := h.Peek().(*Task)
		pe.Name = pe.Name + "_p" + strconv.Itoa(cnt)
		e := h.Poll()
		if t, ok := e.(*Task); ok {
			fmt.Println(t)
		}
	}
}
