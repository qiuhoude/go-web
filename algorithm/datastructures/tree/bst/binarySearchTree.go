package bst

import (
	"container/list"
	"fmt"
	"github.com/qiuhoude/go-web/algorithm/datastructures/stack"
	"strings"
)

// 可比较的接口
type Comparable interface {
	// 比较大小 相等返回0 , 当前这个数小返回负数 ,当前数大返回正数
	CompareTo(o Comparable) int
}

type node struct {
	val         Comparable // 包含的值
	left, right *node      // 左右节点
	size        int        // 子节点的数量
	depth       int        // 该节点的深度
}

type BST struct {
	root *node // 根节点
}

func (t *BST) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

func (t *BST) IsEmpty() bool {
	return t.Size() == 0
}

// 向二分搜索树中添加新的元素
func (t *BST) Add(e Comparable) {
	t.root, _ = add(t.root, e, 0)
}

// 第二bool值表示true表示添加成功, false表示添加失败
func add(n *node, e Comparable, depth int) (*node, bool) {
	if n == nil {
		rn := &node{val: e, depth: depth}
		rn.size++
		return rn, true
	}
	var addSuc bool
	if e.CompareTo(n.val) < 0 { // e值小，挂载到左边
		n.left, addSuc = add(n.left, e, depth+1)
	} else if e.CompareTo(n.val) > 0 { // e值大,挂载到右边
		n.right, addSuc = add(n.right, e, depth+1)
	} else { //相等的情况先不做处理
	}
	if addSuc { // 添加成功让 size++
		n.size++
	}
	return n, addSuc
}

// 查看是否包含此元素e
func (t *BST) Contains(e Comparable) bool {
	return contains(t.root, e)
}

// 看以node为根的二分搜索树中是否包含元素e, 递归算法
func contains(n *node, e Comparable) bool {
	if n == nil {
		return false
	}
	if e.CompareTo(n.val) < 0 {
		return contains(n.left, e)
	} else if e.CompareTo(n.val) > 0 {
		return contains(n.right, e)
	} else { //e.CompareTo(n.val) == 0
		return true
	}
}

type TraverseFunc func(e Comparable)

// 前序遍历,最常见的变量方式 preOrder traverse
func (t *BST) PreOrder(f TraverseFunc) {
	preOrder(t.root, f)
}

func preOrder(n *node, f TraverseFunc) {
	if n == nil {
		return
	}
	f(n.val)
	preOrder(n.left, f)
	preOrder(n.right, f)
}

// 中序遍历,是从小到大
func (t *BST) InOrder(f TraverseFunc) {
	inOrder(t.root, f)
}

func inOrder(n *node, f TraverseFunc) {
	if n == nil {
		return
	}
	inOrder(n.left, f)
	f(n.val)
	inOrder(n.right, f)
}

// 后续序遍历
func (t *BST) PostOrder(f TraverseFunc) {
	postOrder(t.root, f)
}

func postOrder(n *node, f TraverseFunc) {
	if n == nil {
		return
	}
	postOrder(n.left, f)
	postOrder(n.right, f)
	f(n.val)
}

//前序遍历 非递归的方式 NR=non recursion
func (t *BST) PreOrderNR(f TraverseFunc) {
	n := t.root
	if n == nil {
		return
	}
	s := stack.New()
	s.Push(n)
	for !s.IsEmpty() {
		tn, _ := s.Pop().(*node)
		f(tn.val)
		if tn.right != nil {
			s.Push(tn.right)
		}
		if tn.left != nil {
			s.Push(tn.left)
		}
	}
}

//前序遍历 非递归的方式2
func (t *BST) PreOrderNR2(f TraverseFunc) {
	n := t.root
	if n == nil {
		return
	}
	s := stack.New()
	for n != nil || !s.IsEmpty() {
		for n != nil {
			f(n.val)
			s.Push(n)
			n = n.left
		}
		n = s.Pop().(*node)
		n = n.left
	}
}

// 中序遍历非递归
func (t *BST) InOrderNR(f TraverseFunc) {
	n := t.root
	if n == nil {
		return
	}
	s := stack.New()
	for n != nil || !s.IsEmpty() {
		for n != nil { // 找到最小的
			s.Push(n)
			n = n.left
		}
		n, _ = s.Pop().(*node)
		f(n.val)
		n = n.right
	}
}

//中序遍历非递归的第2种方式
func (t *BST) InOrderNR2(f TraverseFunc) {
	n := t.root
	if n == nil {
		return
	}
	s := stack.New()
	for n != nil || !s.IsEmpty() {
		if n != nil {
			s.Push(n)
			n = n.left
		} else {
			n = s.Pop().(*node)
			f(n.val)
			n = n.right
		}
	}
}

// 后续遍历 非递归方式
func (t *BST) PostOrderNR(f TraverseFunc) {
	cur := t.root
	if cur == nil {
		return
	}
	s := stack.New()
	var pre *node //Using a pre pointer to record the last visted node
	for cur != nil || !s.IsEmpty() {
		for cur != nil {
			s.Push(cur)
			cur = cur.left
		}
		cur, _ = s.Pop().(*node)
		if cur.right == nil || pre == cur.right {
			f(cur.val)
			pre = cur
			cur = nil
		} else {
			s.Push(cur)
			cur = cur.right
		}
	}
}

// Morris 方式进行遍历搜索树,不借用占 只用普通的变量
func (t *BST) PreOrderMorris(f TraverseFunc) {
	cur := t.root
	if cur == nil {
		return
	}
	for cur != nil {
		if cur.left == nil {
			f(cur.val)
			cur = cur.right
		} else {
			prev := cur.left
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}
			if prev.right == nil {
				f(cur.val)
				prev.right = cur
				cur = cur.left
			} else {
				prev.right = nil
				cur = cur.right
			}
		}
	}
}

func (t *BST) InOrderMorris(f TraverseFunc) {
	cur := t.root
	if cur == nil {
		return
	}
	for cur != nil {
		if cur.left == nil {
			f(cur.val)
			cur = cur.right
		} else {
			prev := cur.left
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}
			if prev.right == nil {
				prev.right = cur
				cur = cur.left
			} else {
				prev.right = nil
				f(cur.val)
				cur = cur.right
			}
		}
	}
}

// 二分搜索树的层序遍历,也是广度遍历,借助队列的结构遍历
func (t *BST) LevelOrder(f TraverseFunc) {
	if t.root == nil {
		return
	}
	l := list.New()
	l.PushBack(t.root)
	for l.Len() != 0 {
		n, _ := l.Remove(l.Front()).(*node)
		f(n.val)
		if n.left != nil {
			l.PushBack(n.left)
		}
		if n.right != nil {
			l.PushBack(n.right)
		}
	}
}

// 该树的最大深度
func (t *BST) MaxDepth() int {
	return maxDepth(t.root)

	//if t.root == nil {
	//	return 0
	//}
	//maxD := 0
	//calcDepth(t.root, 0, &maxD)
	//return maxD

}
func calcDepth(n *node, depth int, maxDepth *int) {
	if n == nil {
		return
	}
	if *maxDepth < depth+1 {
		*maxDepth = depth + 1
	}
	calcDepth(n.left, depth+1, maxDepth)
	calcDepth(n.right, depth+1, maxDepth)
}

func maxDepth(n *node) int {
	if n == nil {
		return 0
	}
	// 可用理解为是后序遍历
	// 在跟父点处比较大小
	leftDepth := maxDepth(n.left)
	rightDepth := maxDepth(n.right)
	return max(leftDepth, rightDepth) + 1
}

func minDepth(n *node) int {
	if n == nil {
		return 0
	}
	if n.left == nil { // 若左子树为空，则右子树的深度为为该节点的深度
		return minDepth(n.right) + 1
	}
	if n.right == nil {
		return minDepth(n.left) + 1
	}
	leftDepth := minDepth(n.left)
	rightDepth := minDepth(n.right)
	return min(leftDepth, rightDepth) + 1
}

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

// 寻找二分搜索树的最小元素
func (t *BST) Minimum() Comparable {
	if t.root == nil {
		return nil
	}
	return minimum(t.root).val
}

func minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最小元素 非递归方式
func (t *BST) MinimumNR() Comparable {
	if t.root == nil {
		return nil
	}
	tn := t.root
	for {
		if tn.left == nil {
			break
		}
		tn = tn.left
	}
	return tn.val
}

// 寻找二分搜索树的最大元素
func (t *BST) Maximum() Comparable {
	if t.root == nil {
		return nil
	}
	return maximum(t.root).val
}

func maximum(n *node) *node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

func (t *BST) MaximumNR() Comparable {
	if t.root == nil {
		return nil
	}
	tn := t.root
	for {
		if tn.right == nil {
			break
		}
		tn = tn.right
	}
	return tn.val
}

// 从二分搜索树中删除最小值所在节点, 返回最小值
func (t *BST) RemoveMin() Comparable {
	if t.root == nil {
		return nil
	}
	ret := minimum(t.root)
	t.root = removeMin(t.root)
	return ret.val
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根 和 是否删除成功
func removeMin(n *node) *node {
	if n.left == nil {
		// 将要删除的右节点挂载父节点上,通过返回值返给服节点
		tn := n.right
		n.right = nil // 置空 gc回收
		opDepth(tn, -1)
		return tn
	}
	n.left = removeMin(n.left)
	n.size--
	return n
}

// 对节点下面的所有节点进行操作
func opDepth(n *node, op int) {
	if n == nil {
		return
	}
	n.depth = n.depth + op
	opDepth(n.left, op)
	opDepth(n.right, op)
}

// 从二分搜索树中删除最大值所在节点
func (t *BST) RemoveMax() Comparable {
	if t.root == nil {
		return nil
	}
	ret := maximum(t.root)
	t.root = removeMax(t.root)
	return ret.val
}

func removeMax(n *node) *node {
	if n.right == nil {
		tn := n.left
		n.left = nil
		opDepth(tn, -1)
		return tn
	}
	n.right = removeMax(n.right)
	n.size--
	return n
}

// 移除指定元素,返回true表示移除成功
func (t *BST) Remove(c Comparable) bool {
	if !t.Contains(c) { // 不存在删除失败
		return false
	}
	n := remove(t.root, c)
	t.root = n
	return true
}

// 移除对应元素,返回跟节点
func remove(n *node, c Comparable) *node {
	if n == nil {
		return nil
	}
	if c.CompareTo(n.val) < 0 {
		n.left = remove(n.left, c)
		if n.left != nil {
			n.size--
		}
		return n
	} else if c.CompareTo(n.val) > 0 {
		n.right = remove(n.right, c)
		if n.right != nil {
			n.size--
		}
		return n
	} else { //相等,进行移除
		// 以下分几种情况

		// 1. 待删除节点左子树为空的情况
		if n.left == nil { // 将右子树的数据反给上层
			tn := n.right
			n.left = nil
			opDepth(tn, -1)
			return tn
		}
		// 2. 待删除节点右子树为空的情况
		if n.right == nil {
			tn := n.left
			n.left = nil
			opDepth(tn, -1)
			return tn
		}

		// 3. 左右都有数据,
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := minimum(n.right)
		successor.right = removeMin(n.right)
		successor.left = n.left
		successor.size = n.size - 1
		successor.depth = n.depth
		n.left = nil
		n.right = nil
		return successor
	}
}

func (t *BST) String() string {
	var sb strings.Builder
	//generateBSTString(t.root, 0, &sb)
	generateBSTLevelString(t.root, &sb)
	return sb.String()
}

func generateBSTLevelString(n *node, sb *strings.Builder) {
	l := list.New()
	l.PushBack(n)
	//curDepth := 0

	curNodeCnt := 1
	nextNodeCnt := 0
	for l.Len() != 0 {
		n, _ := l.Remove(l.Front()).(*node)
		//if n.depth > curDepth {
		//	curDepth = n.depth
		//	sb.WriteRune('\n')
		//}

		sb.WriteString(fmt.Sprintf("%v ", n.val))
		curNodeCnt--
		if n.left != nil {
			l.PushBack(n.left)
			nextNodeCnt++
		}
		if n.right != nil {
			l.PushBack(n.right)
			nextNodeCnt++
		}
		if curNodeCnt == 0 {
			sb.WriteRune('\n')
			curNodeCnt = nextNodeCnt
			nextNodeCnt = 0
		}
	}
}

func generateBSTString(n *node, depth int, sb *strings.Builder) {
	if n == nil {
		//generateDepthString(depth, sb)
		//sb.WriteString("\n")
		return
	}
	generateDepthString(depth, sb)
	sb.WriteString(fmt.Sprintf("val:%v,child:%d,dh:%d\n", n.val, n.size-1, n.depth))
	generateBSTString(n.left, depth+1, sb)
	generateBSTString(n.right, depth+1, sb)
}

func generateDepthString(depth int, sb *strings.Builder) {
	for i := 0; i < depth; i++ {
		sb.WriteString("--")
	}
}
