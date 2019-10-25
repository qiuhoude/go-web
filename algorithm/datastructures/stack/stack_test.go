package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New()
	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Pop()

	t.Log(s)
}

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	//s.Push(5)
	//s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Pop()
	s.Pop()
	t.Log(s)
	t.Log(s.Peek())
}

// 双栈的的应用
// 1. 浏览器前一页 后一页
// 2. 命令设计模式中的 undo redo
// 3. 四则运算 2+3*9-1 这种, 数字和运算符号,各用一个栈,
// 如果比运算符栈顶元素的优先级高，就将当前运算符压入栈；如果比运算符栈顶元素的优先级低
//	或者相同，从运算符栈中取栈顶运算符，从操作数栈的栈顶取 2 个操作数，然后进行计算，再
//	把计算完的结果压入操作数栈，继续比较

type Browser struct {
	back, forward IStack
}

func NewBrowser() *Browser {
	return &Browser{
		back:    NewArrayStack(),
		forward: New(),
	}
}

func (b *Browser) CanBack() bool {
	return !b.back.IsEmpty()
}

func (b *Browser) CanForward() bool {
	return !b.forward.IsEmpty()
}
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forward.Flush()
}

func (b *Browser) Forward() {
	if b.CanForward() {
		top := b.forward.Pop()
		b.back.Push(top)
		fmt.Printf("forward to %+v\n", top)
	}
}
func (b *Browser) Back() {
	if b.CanBack() {
		top := b.back.Pop()
		b.forward.Push(top)
		fmt.Printf("back to %+v\n", top)
	}
}

func (b *Browser) PushBack(addr string) {
	b.back.Push(addr)
}
func TestStack_Browser(t *testing.T) {
	b := NewBrowser()
	b.PushBack("www.qq.com")
	b.PushBack("www.baidu.com")
	b.PushBack("www.sina.com")
	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.taobao.com")
	if b.CanForward() {
		b.Forward()
	}
}
