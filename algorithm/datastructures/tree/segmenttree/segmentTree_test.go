package segmenttree

import (
	"fmt"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	d := []int{-2, 0, 3, -5, 2, -1}

	data := make([]interface{}, 0, len(d))
	for _, v := range d {
		data = append(data, v)
	}
	tree := NewSegmentTree(data, func(l, r interface{}) interface{} {
		nl := l.(int)
		nr := r.(int)
		return max(nl, nr)
	})
	fmt.Println(tree.Query(0, 2))
	//fmt.Println(tree)
	tree.Set(1, 100)
	fmt.Println("---------------set after-------------")
	fmt.Println(tree.Query(0, 2))
	//fmt.Println(tree)
	fmt.Println(tree.Query(2, 5))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestSegmentTree_Depth(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(i, depth(i))
	}
}
