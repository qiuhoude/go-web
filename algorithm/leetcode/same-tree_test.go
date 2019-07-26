package leetcode

// 100. 相同的树
// https://leetcode-cn.com/problems/same-tree/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 都是nil
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) { //只有一个其中一个nil
		return false
	}
	return q.Val == p.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
