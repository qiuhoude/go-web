package queue

import (
	"fmt"
	"strconv"
	"strings"
)

// 环形数组队列
type CircularQueue struct {
	data                 []interface{}
	head, tail, capacity int
}

func NewCircularQueue(cap int) *CircularQueue {
	rcap := cap + 1
	return &CircularQueue{
		data:     make([]interface{}, rcap, rcap),
		head:     0,
		tail:     0,
		capacity: rcap,
	}
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % q.capacity
	return ret
}

func (q *CircularQueue) Enqueue(v interface{}) bool {
	if (q.tail+1)%q.capacity == q.head { //满了
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircularQueue) Len() int {
	// 注意此时getSize的逻辑:
	// 如果tail >= head，非常简单，队列中的元素个数就是tail - head
	// 如果tail < head，说明我们的循环队列"循环"起来了，此时，队列中的元素个数为：
	// tail  + capacity - head;tail  + capacity 表示绕过了环
	if q.tail >= q.head {
		return q.tail - q.head
	} else {
		return q.tail + q.capacity - q.head
	}
}

func (q *CircularQueue) String() string {
	var sb strings.Builder
	sb.WriteString("head ")
	for i := q.head; i%q.capacity != q.tail; i++ {
		sb.WriteString(fmt.Sprintf("%v", q.data[i%q.capacity]))
		if (i+1)%q.capacity != q.tail { // 非倒数第一个
			sb.WriteString(",")
		}
	}
	sb.WriteString(" tail size:" + strconv.Itoa(q.Len()))
	sb.WriteString(", cap:" + strconv.Itoa(q.capacity))
	return sb.String()
}

// 动态扩容和缩容
func (q *CircularQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % q.capacity

	len := q.Len()
	// 长度等于容量的 1/4 进行缩容 容量的 1/2
	if len == q.capacity/4 && q.capacity/2 != 0 {
		q.resize(q.capacity / 2)
	}
	return ret
}

func (q *CircularQueue) Offer(v interface{}) bool {
	if (q.tail+1)%q.capacity == q.head { // 满了
		q.resize(2 * q.capacity) // 进行2倍的扩容
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// 动态调整
func (q *CircularQueue) resize(cap int) {
	rcap := cap + 1
	nData := make([]interface{}, rcap, rcap)
	len := q.Len()
	for i := 0; i < len; i++ {
		nData[i] = q.data[(q.head+i)%q.capacity]
	}
	// 重新赋值
	q.head = 0
	q.tail = len
	q.data = nData
	q.capacity = rcap
}
