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
