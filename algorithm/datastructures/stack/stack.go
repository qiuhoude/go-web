// 栈结构
package stack

type (
	Stack struct {
		top    *node // 栈顶
		length int
	}

	node struct {
		value interface{}
		pre   *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.pre
	s.length--
	return n.value
}

func (s *Stack) Push(value interface{}) {
	n := s.top
	s.top = &node{value, n}
	s.length++
}
