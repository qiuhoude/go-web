package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// 257. 二叉树的所有路径
// https://leetcode-cn.com/problems/binary-tree-paths/

func TestB(t *testing.T) {
	lr := TreeNode{Val: 5}
	l := TreeNode{Val: 2, Right: &lr}
	r := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &l, Right: &r}
	fmt.Println(binaryTreePaths(&root))
}

// 递归写法
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	} else {
		var ret []string
		if root.Left != nil {
			c := strconv.Itoa(root.Val) + "->"
			sub := binaryTreePaths(root.Left)
			for _, v := range sub {
				ret = append(ret, c+v)
			}
		}
		if root.Right != nil {
			c := strconv.Itoa(root.Val) + "->"
			sub := binaryTreePaths(root.Right)
			for _, v := range sub {
				ret = append(ret, c+v)
			}
		}
		return ret
	}
}
