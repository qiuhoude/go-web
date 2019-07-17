package stack

import "fmt"

func Example() {
	//创建一个stack
	s := New()
	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	p := s.Pop()
	fmt.Println(p)
	fmt.Println(s.Len())

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}

	//Output:
	//2
	//3
	//3
	//4
	//5
}
