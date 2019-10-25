package queue

import (
	"fmt"
	"strings"
)

type (
	node struct {
		data interface{}
		next *node
	}

	LinkedQueue struct {
		head *node
		tail *node
		len  int
	}
)

func newNode(v interface{}) *node {
	return &node{v, nil}
}

var emptyNode = newNode(nil)

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil, nil, 0}
}

func (q *LinkedQueue) Enqueue(v interface{}) bool {
	nn := newNode(v)
	if q.head == nil {
		q.head = nn
	}
	if q.tail != nil {
		q.tail.next = nn
	}
	q.tail = nn
	q.len++
	return true
}

func (q *LinkedQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.head
	q.head = q.head.next
	q.len--
	return ret.data
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *LinkedQueue) Len() int {
	return q.len
}

func (q *LinkedQueue) String() string {
	var sb strings.Builder
	sb.WriteString("head ")
	cur := q.head
	for ; cur != nil; cur = cur.next {
		sb.WriteString(fmt.Sprintf("%v->", cur.data))
	}
	sb.WriteString("tail")
	return sb.String()
}
