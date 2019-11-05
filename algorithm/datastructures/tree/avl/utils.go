package avl

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 判断该二叉树是否是一棵二分搜索树
func IsBST(tree *AVLTree) bool {
	if tree == nil {
		return false
	}
	var ls []interface{}
	tree.InOrder(func(i interface{}) {
		ls = append(ls, i)
	})
	for i := 1; i < len(ls); i++ {
		cmp := tree.compareFunc(ls[i-1], ls[i])
		if cmp > 0 {
			return false
		}
	}
	return true
}

// 判断该二叉树是否是一棵平衡二叉树
func IsBalanced(tree *AVLTree) bool {
	if tree == nil {
		return false
	}
	return isBalanced(tree.root)
}

//判断以Node为根的二叉树是否是一棵平衡二叉树，递归算法
func isBalanced(n *avlNode) bool {
	if n == nil {
		return true
	}
	balanceFactor := getBalanceFactor(n)
	if abs(balanceFactor) > 1 {
		return false
	}
	// 左子树右子树都是平衡二叉树才是
	return isBalanced(n.left) && isBalanced(n.right)
}
