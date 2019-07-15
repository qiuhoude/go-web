package tree

import (
	"fmt"
	"strconv"
	"testing"
)

type Integer int

func (i Integer) CompareTo(o Comparable) int {
	oi, ok := o.(Integer)
	if ok {
		return int(i - oi) // 简单这么写,没有考虑溢出的情况
	}
	return -1
}
func (i Integer) String() string {
	return strconv.Itoa(int(i))
}

func TestBST(t *testing.T) {
	bst := BST{}
	nums := []Integer{Integer(5), Integer(3), Integer(6), Integer(8), Integer(4), Integer(2)}
	for i := 0; i < len(nums); i++ {
		bst.Add(nums[i])
	}

	bst.PreOrder(func(e Comparable) {
		fmt.Println(e)
	})

	fmt.Println(bst)
}
