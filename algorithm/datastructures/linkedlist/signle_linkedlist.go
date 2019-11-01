package linkedlist

import (
	"fmt"
	"strings"
)

/*
单链表
*/

type ListNode struct {
	next  *ListNode
	value interface{}
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func (this *ListNode) String() string {
	return fmt.Sprintf("%v", this.value)
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(nil), 0}
}

func (this *LinkedList) Len() uint {
	return this.length
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for ; cur != nil; cur = cur.next {
		if cur == p {
			break
		}
		pre = cur
	}
	if cur == nil { // 没有找到
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 反转链表
func (this *LinkedList) Reverse() {
	cur := this.head.next
	var pre *ListNode = nil
	var tmp *ListNode = nil
	for cur != nil {
		tmp = pre
		pre = cur
		cur = cur.next
		pre.next = tmp
	}
	/*for ; cur != nil; {
		tmp = cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}*/
	this.head.next = pre
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for ; cur.next != nil; cur = cur.next {
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head
	for i := uint(0); cur.next != nil; i++ {
		cur = cur.next
		if i == index {
			return cur
		}
	}
	return nil
}

func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head
	for ; cur != nil; cur = cur.next {
		if cur == p {
			break
		}
		pre = cur
	}
	if cur == nil {
		return false
	}

	pre.next = cur.next
	cur.next = nil
	this.length--
	return true
}

// 判断单链表是否有环
// 思路: 快 慢 指针,不断往后移动,只要快慢指针指向同位置说明有环
func (this *LinkedList) HasCycle() bool {
	if nil != this.head {
		slow := this.head
		fast := this.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false

}

// 找到中间节点 当有 len值时 直接 len>0; index=len/2
func (this *LinkedList) FindMidNode2() *ListNode {
	if this.length <= 0 {
		return nil
	}
	index := this.length / 2
	return this.FindByIndex(uint(index))
}

// 找到中间节点
func (this *LinkedList) FindMidNode() *ListNode {
	if this.head.next == nil {
		return nil
	}
	fast := this.head
	slow := this.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	if fast != nil {
		slow = slow.next
	}
	return slow
}

func (this *LinkedList) String() string {
	var sb strings.Builder
	cur := this.head.next
	sb.WriteString("head ")
	for ; cur != nil; cur = cur.next {
		_, _ = fmt.Fprintf(&sb, "->%v", cur.GetValue())
	}
	sb.WriteString(" tail")
	return sb.String()
}

/*
两个有序单链表合并, 归并排序的链表结构可以此
*/
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil && l2 == nil {
		return nil
	}
	nL := NewLinkedList()
	cur := nL.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			cur.next = cur2
			cur2 = cur2.next
		} else {
			cur.next = cur1
			cur1 = cur1.next
		}
		cur = cur.next
	}
	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil {
		cur.next = cur2
	}

	return nL
}

// 拷贝值的方式 重新创建新列表
func MergeSortedList2(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil && l2 == nil {
		return nil
	}
	nL := NewLinkedList()

	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) > cur2.value.(int) {
			nL.InsertToTail(cur2.value)
			cur2 = cur2.next
		} else {
			nL.InsertToTail(cur1.value)
			cur1 = cur1.next
		}
	}
	for cur1 != nil || cur2 != nil {
		if cur1 != nil {
			nL.InsertToTail(cur1.value)
			cur1 = cur1.next
		}
		if cur2 != nil {
			nL.InsertToTail(cur2.value)
			cur2 = cur2.next
		}
	}
	return nL
}

/*
删除倒数第N个节点
思路: 还是快 慢 双指针, 让快指针先走n部,然后快慢指针同时走, 快指针到尾部就结束
此时慢指针下一个node就是要目标node
*/
func (this *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == this.head || nil == this.head.next {
		return
	}
	// 相聚
	fast := this.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if nil == fast {
		return
	}

	slow := this.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
