package leetcode

//437. 路径总和 III
//https://leetcode-cn.com/problems/path-sum-iii/

// 遍历每个节点为起点, 调用一个子过程找到路径和
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := findNodeSum(root, sum)
	cnt += pathSum(root.Left, sum)
	cnt += pathSum(root.Right, sum)
	return cnt
}

func findNodeSum(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}

	ret := 0
	if node.Val == num {
		ret += 1 // 找到了+1
	}
	ret += findNodeSum(node.Left, num-node.Val)
	ret += findNodeSum(node.Right, num-node.Val)
	return ret
}
