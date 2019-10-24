package linkedlist

import "testing"

func TestLinkedList_InsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i <= 10; i++ {
		l.InsertToHead(i)
	}
	t.Log(l)
}

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}
	t.Log(l)
}

func TestLinkedList_FindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}

	t.Log(l.FindByIndex(uint(0)))
	t.Log(l.FindByIndex(uint(5)))
	t.Log(l.FindByIndex(uint(10)))
	t.Log(l.FindByIndex(uint(300)))
}

func TestLinkedList_DeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 1; i <= 10; i++ {
		l.InsertToTail(i)
	}
	t.Log(l)

	t.Log(l.DeleteNode(l.FindByIndex(3)))
	t.Log(l)

	t.Log(l.DeleteNode(l.FindByIndex(6)))
	t.Log(l)

	t.Log(l.DeleteNode(l.FindByIndex(10)))
	t.Log(l)
}

func TestLinkedList_Reverse(t *testing.T) {
	t.Log(l)
	l.Reverse()
	t.Log(l)

	t.Log(emptyL)
	emptyL.Reverse()
	t.Log(emptyL)

}

var l, emptyL, oneL, oddL, evenL *LinkedList

func init() {
	emptyL = NewLinkedList()

	n5 := &ListNode{value: 5}
	n4 := &ListNode{value: 4, next: n5}
	n3 := &ListNode{value: 3, next: n4}
	n2 := &ListNode{value: 2, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l = &LinkedList{head: &ListNode{next: n1}}

	oneL = NewLinkedList()
	oneL.InsertToTail(1)

	oddL = NewLinkedList()
	for i := 1; i <= 9; i++ {
		oddL.InsertToTail(i)
	}
	evenL = NewLinkedList()
	for i := 1; i <= 10; i++ {
		evenL.InsertToTail(i)
	}

}

func TestLinkedList_FindMidNode(t *testing.T) {

	tests := []struct {
		name string
		l    *LinkedList
		want interface{}
	}{
		{"空链表 => ", emptyL, nil},
		{"一个元素 => ", oneL, 1},
		{"奇数个元素 => ", oddL, 5},
		{"偶数个元素 => ", evenL, 6},
	}
	for _, tt := range tests {
		ts := tt
		t.Run(tt.name, func(t *testing.T) {
			node := ts.l.FindMidNode()
			/*
				var got interface{} = nil
				if node != nil {
					got = node.value
				}
				if got != tt.want {
					t.Errorf("FindMidNode() = %v, want %v", got, ts.want)
				}*/

			node2 := ts.l.FindMidNode2()
			if node != node2 {
				t.Errorf("FindMidNode() = %v, want %v", node, node2)
			}
		})
	}
}

func TestLinkedList_HasCycle(t *testing.T) {
	n6 := &ListNode{value: 6}
	n5 := &ListNode{value: 5, next: n6}
	n4 := &ListNode{value: 4, next: n5}
	n3 := &ListNode{value: 3, next: n4}
	n2 := &ListNode{value: 2, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	n6.next = n5
	ll := &LinkedList{head: &ListNode{next: n1}}
	t.Log(ll.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &ListNode{value: 9}
	n4 := &ListNode{value: 7, next: n5}
	n3 := &ListNode{value: 5, next: n4}
	n2 := &ListNode{value: 3, next: n3}
	n1 := &ListNode{value: 1, next: n2}
	l1 := &LinkedList{head: &ListNode{next: n1}}

	n12 := &ListNode{value: 12}
	n11 := &ListNode{value: 11, next: n12}
	n10 := &ListNode{value: 10, next: n11}
	n9 := &ListNode{value: 8, next: n10}
	n8 := &ListNode{value: 6, next: n9}
	n7 := &ListNode{value: 4, next: n8}
	n6 := &ListNode{value: 2, next: n7}
	l2 := &LinkedList{head: &ListNode{next: n6}}

	t.Log(l1)
	t.Log(l2)
	nl := MergeSortedList2(l1, l2)
	t.Log("merge after ", nl)
}

func TestDeleteBottomN(t *testing.T) {
	t.Log(l)
	l.DeleteBottomN(1)
	t.Log(l)
}
