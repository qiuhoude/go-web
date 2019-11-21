package queue

import "github.com/qiuhoude/go-web/algorithm/datastructures/minheap"

// 优先级队列
type PriorityQueue struct {
	heap *minheap.Heap
}

func NewPriorityQueue(f minheap.CompareFunc) *PriorityQueue {
	return &PriorityQueue{heap: minheap.NewHeap(f)}
}

func (q *PriorityQueue) Enqueue(v interface{}) bool {
	q.heap.Add(v)
	return true
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.heap.Poll()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.heap.Len() == 0
}

func (q *PriorityQueue) Len() int {
	return q.heap.Len()
}

// 移除对应的值
func (q *PriorityQueue) Remove(e interface{}, eqFunc func(e, b interface{}) bool) interface{} {
	return q.heap.Remove(e, eqFunc)
}
