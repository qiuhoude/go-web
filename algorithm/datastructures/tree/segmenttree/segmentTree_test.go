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
		return nl + nr
	})
	fmt.Println(tree.Query(0, 2))
	fmt.Println(tree.Query(2, 5))
	fmt.Println(tree)
}

func TestSegmentTree_Depth(t *testing.T) {

}
