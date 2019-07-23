package bst

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
	fmt.Println("PreOrder-------------------")
	bst.PreOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("InOrder-------------------")
	bst.InOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("PostOrder-------------------")
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
	fmt.Println("-------------------")
	bst.PreOrderMorris(func(e Comparable) {
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
	fmt.Println("-------------------")
	bst.InOrderNR2(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.InOrderMorris(func(e Comparable) {
		fmt.Println(e)
	})
}

func TestBST_PostOrderNR(t *testing.T) {
	bst := generateBST()
	bst.PostOrder(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.PostOrderNR(func(e Comparable) {
		fmt.Println(e)
	})
}

func TestBST_LevelOrder(t *testing.T) {
	bst := generateBST()
	bst.PreOrderNR(func(e Comparable) {
		fmt.Println(e)
	})
	fmt.Println("-------------------")
	bst.LevelOrder(func(e Comparable) {
		fmt.Println(e)
	})
}

func TestBst_MaxDepth(t *testing.T) {
	bst := generateBST()
	t.Log("最大深度 ", bst.MaxDepth())
}

func TestBST_Minimum(t *testing.T) {
	bst := generateBST()
	min := bst.Minimum().(Integer)
	minNR := bst.MinimumNR().(Integer)
	t.Log("最小值: ", min, minNR)

}

func TestBST_Maximum(t *testing.T) {
	bst := generateBST()
	max := bst.Maximum().(Integer)
	t.Log("最大值: ", max)
}

func TestBST_RemoveMin(t *testing.T) {
	bst := generateBST()
	for bst.Size() > 0 {
		m := bst.RemoveMin().(Integer)
		fmt.Println("移除最小值: ", m)
		fmt.Println(bst)
	}
}

func TestBST_RemoveMax(t *testing.T) {
	bst := generateBST()
	for !bst.IsEmpty() {
		m := bst.RemoveMax().(Integer)
		fmt.Println("移除大小值: ", m)
		fmt.Println(bst)
	}
}

func TestBST_Remove(t *testing.T) {
	bst := generateBST()
	b := bst.Remove(Integer(8))
	if b {
		fmt.Println("删除成功")
		fmt.Println(bst)
	} else {
		fmt.Println("删除失败")
	}

}
