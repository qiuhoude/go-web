package stack

import (
	"fmt"
	"strings"
)

// 数组站
type ArrayStack struct {
	top  int
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{0, make([]interface{}, 0, 8)}
}

func (s *ArrayStack) Len() int {
	return s.top
}

func (s *ArrayStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top-1]
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := s.data[s.top-1]
	s.data[s.top-1] = nil
	s.top--
	return ret
}

func (s *ArrayStack) ToSlice() []interface{} {
	ret := make([]interface{}, s.top)
	j := 0
	for i := s.top - 1; i >= 0; i-- {
		ret[j] = s.data[i]
		j++
	}
	return ret
}
func (s *ArrayStack) Push(value interface{}) {
	if s.top >= len(s.data) {
		s.data = append(s.data, value)
	} else {
		s.data[s.top] = value
	}
	s.top++
}
func (s *ArrayStack) Flush() {
	s.top = 0
	for i := range s.data {
		s.data[i] = nil
	}
}
func (s *ArrayStack) String() string {
	sb := new(strings.Builder)
	sb.WriteString("top ")
	for i := s.top - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%v ", s.data[i]))
	}
	return sb.String()
}
