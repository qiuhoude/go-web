package queue

type Queue interface {
	// 入队列
	Enqueue(v interface{}) bool
	// 出队列
	Dequeue() interface{}

	IsEmpty() bool

	Len() int
}
