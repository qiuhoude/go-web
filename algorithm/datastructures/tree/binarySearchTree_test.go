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

func generateBST() *BST {
	/////////////////
	//      5      //
	//    /   \    //
	//   3    6    //
	//  / \    \   //
	// 2  4     8  //
	/////////////////
	bst := &BST{}
	nums := []Integer{Integer(5), Integer(3), Integer(6), Integer(8), Integer(4), Integer(2), Integer(7)}
	for i := 0; i < len(nums); i++ {
		bst.Add(nums[i])
	}
	return bst
}

func TestBST_traverse(t *testing.T) {
	bst := generateBST()
	bst.PreOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.InOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.PostOrder(func(e Comparable) {
		fmt.Println(e)
	})
}

func TestBST_PreOrderNR(t *testing.T) {
	bst := generateBST()
	bst.PreOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.PreOrderNR(func(e Comparable) {
		fmt.Println(e)
	})
}

func TestBST_InOrderNR(t *testing.T) {
	bst := generateBST()
	bst.InOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.InOrderNR(func(e Comparable) {
		fmt.Println(e)
	})
}
