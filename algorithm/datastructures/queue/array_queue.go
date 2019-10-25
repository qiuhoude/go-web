package queue

import (
	"fmt"
	"strings"
)

// 数组队列
type ArrayQueue struct {
	data     []interface{}
	head     int
	tail     int
	capacity int
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, cap, cap),
		head:     0,
		tail:     0,
		capacity: cap,
	}
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[q.head]
	q.data[q.head] = nil
	q.head++
	return ret
}

func (q *ArrayQueue) Enqueue(v interface{}) bool {
	len := q.Len()
	if len >= q.capacity { //满了
		return false
	}
	if q.tail >= q.capacity { // 进行搬迁
		for i := 0; i < len; i++ {
			q.data[i] = q.data[i+q.head]
			q.data[i+q.head] = nil
		}
		q.head = 0
		q.tail = len
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.tail == q.head
}

func (q *ArrayQueue) Len() int {
	return q.tail - q.head
}

func (q *ArrayQueue) String() string {
	var sb strings.Builder
	sb.WriteString("head ")
	for i := q.head; i < q.tail; i++ {
		sb.WriteString(fmt.Sprintf("%v ", q.data[i]))
	}
	sb.WriteString("tail")
	return sb.String()
}
