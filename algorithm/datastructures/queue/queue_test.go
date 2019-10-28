package queue

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
}

func TestArrayQueue_Dequeue(t *testing.T) {
	q := NewArrayQueue(5)
	for i := 1; i < 10; i++ {
		if q.Enqueue(i) {
			t.Log(q)
		}
	}
	for i := 0; i < 10; i++ {
		q.Dequeue()
		t.Log(q)
	}

}

func TestLinkedQueue_Dequeue(t *testing.T) {
	q := NewLinkedQueue()
	for i := 1; i < 10; i++ {
		if q.Enqueue(i) {
			t.Log(q.Len())
			t.Log(q)
		}
	}
	for i := 0; i < 10; i++ {
		v := q.Dequeue()
		t.Log("v =", v)
		t.Log(q)
	}

}

func TestCircularQueue_Enqueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
}

func TestCircularQueue_Dequeue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
	t.Log(q.Dequeue())
	t.Log(q)
	q.Enqueue(6)
	t.Log(q)
	q.Dequeue()
	q.Dequeue()
	t.Log(q)
	q.Enqueue(7)
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
}

func TestCircularQueue_Dynamic(t *testing.T) {
	q := NewCircularQueue(5)
	q.Offer(1)
	q.Offer(2)
	q.Offer(3)
	q.Offer(4)
	q.Offer(5)
	q.Offer(6)
	t.Log(q)
	t.Log(q.Poll())
	t.Log(q)
	q.Offer(6)
	t.Log(q)
	q.Poll()
	q.Poll()
	t.Log(q)
	q.Offer(7)
	t.Log(q)
	q.Poll()
	t.Log(q)
	q.Poll()
	t.Log(q)
	q.Poll()
	t.Log(q)
}
