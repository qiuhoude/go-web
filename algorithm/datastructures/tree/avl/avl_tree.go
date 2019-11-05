package avl

import (
	"container/list"
	"fmt"
	"strings"
)

// 图像化 https://www.cs.usfca.edu/~galles/visualization/Algorithms.html
//

/*
平衡二叉树:
对于任意一个节点,左子树和右子树的高度差不能超过1

平衡二叉树的高度和节点数量之间的关系也是 O(logn)
所以需要 记录高度 和 计算平衡因子
*/

// 节点
type avlNode struct {
	left, right *avlNode
	height      int // 当前节点的高度
	v           interface{}
}

func (n *avlNode) String() string {
	return fmt.Sprintf("h:%v, v:%+v", n.height, n.v)
}

func newAvlNode(v interface{}) *avlNode {
	return &avlNode{
		v:      v,
		height: 1,
	}
}

//获得节点node的高度
func getHeight(n *avlNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

// 获得节点node的平衡因子,左子树高度 - 右子树高度
func getBalanceFactor(n *avlNode) int {
	if n == nil {
		return 0
	}
	return getHeight(n.left) - getHeight(n.right)
}

// avl树
type AVLTree struct {
	root        *avlNode
	size        int
	compareFunc CompareFunc
}

// 比较函数,v表示要操作的的值
type CompareFunc func(v, nodeV interface{}) int

func NewAVLTree(cfunc CompareFunc) *AVLTree {
	return &AVLTree{
		compareFunc: cfunc,
	}
}

func (t *AVLTree) InOrder(f func(interface{})) {
	inOrder(t.root, f)
}

func inOrder(n *avlNode, f func(interface{})) {
	if n == nil {
		return
	}
	inOrder(n.left, f)
	f(n.v)
	inOrder(n.right, f)
}

func (t *AVLTree) Add(v interface{}) {
	t.root, _ = t.add(t.root, v)
}

func (t *AVLTree) add(n *avlNode, v interface{}) (*avlNode, bool) {
	if n == nil {
		t.size++
		return newAvlNode(v), true
	}
	cmp := t.compareFunc(v, n.v)
	var addSuc bool
	if cmp == 0 {
		return n, false
	} else if cmp > 0 {
		n.right, addSuc = t.add(n.right, v)
	} else { //cmp < 0
		n.left, addSuc = t.add(n.left, v)
	}

	// 更新height
	n.height = max(getHeight(n.left), getHeight(n.right)) + 1

	// 计算平衡因子
	balanceFactor := getBalanceFactor(n)

	//if abs(balanceFactor) > 1 { // 平衡因子大1 需要处理
	//	fmt.Println("unbalanced:", balanceFactor)
	//}

	// 维护平衡操作
	if balanceFactor > 1 && getBalanceFactor(n.left) >= 0 { //LL
		// 左边高 右边低
		// LL 型,在父节点的左孩子的左子树添加了新节点，导致根节点的平衡因子变为 +2，二叉树失去平衡
		// 右旋一次即可
		n = t.rightRotate(n)
	} else if balanceFactor < -1 && getBalanceFactor(n.right) <= 0 { //RR
		// 左边低 右边高
		// RR 型 同理
		n = t.leftRotate(n)
	} else if balanceFactor > 1 && getBalanceFactor(n.left) < 0 { //LR
		//LR 就是将新的节点插入到了 n 的左孩子的右子树上导致的不平衡的情况。
		// 这时我们需要的是先对 左孩子进行一次左旋再对 自己 进行一次右旋
		n.left = t.leftRotate(n.left)
		n = t.rightRotate(n)

	} else if balanceFactor < -1 && getBalanceFactor(n.right) > 0 { //RL
		n.right = t.rightRotate(n.right)
		n = t.leftRotate(n)
	}

	return n, addSuc
}

// 计算数高
func treeHeight(n *avlNode) int {
	if n == nil {
		return 0
	}
	return max(treeHeight(n.left), treeHeight(n.right)) + 1
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (t *AVLTree) rightRotate(y *avlNode) *avlNode {
	if y == nil {
		return nil
	}
	x := y.left
	t3 := x.right

	x.right = y
	y.left = t3

	// 更新height
	y.height = max(getHeight(y.left), getHeight(y.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (t *AVLTree) leftRotate(y *avlNode) *avlNode {
	if y == nil {
		return nil
	}
	x := y.right
	t2 := x.left

	x.left = y
	y.right = t2

	// 更新height
	y.height = max(getHeight(y.left), getHeight(y.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1
	return x
}

func (t *AVLTree) findNode(n *avlNode, v interface{}) *avlNode {
	if n == nil {
		return nil
	}
	// 递归方式查找
	cmp := t.compareFunc(v, n.v)
	if cmp > 0 {
		return t.findNode(n.right, v)
	} else if cmp < 0 {
		return t.findNode(n.left, v)
	} else { // ==
		return n
	}

	/*
		//非递归方式
		findN := n
		for findN != nil {
			cmp := t.compareFunc(v, findN.v)
			if cmp > 0 {
				findN = n.right
			} else if cmp < 0 {
				findN = n.left
			} else {
				break
			}
		}
		return findN
	*/
}

func (t *AVLTree) Contains(v interface{}) bool {
	return t.findNode(t.root, v) != nil
}

func (t *AVLTree) Remove(v interface{}) bool {
	if !t.Contains(v) { // 不存在删除失败
		return false
	}
	n := t.remove(t.root, v)
	t.root = n
	return true
}

// 移除对应元素,返回跟节点
func (t *AVLTree) remove(n *avlNode, v interface{}) *avlNode {
	if n == nil {
		return nil
	}
	cmp := t.compareFunc(v, n.v)
	if cmp < 0 {
		n.left = t.remove(n.left, v)
		return n
	} else if cmp > 0 {
		n.right = t.remove(n.right, v)
		return n
	} else { // ==
		// 1. 待删除节点左子树为空的情况
		if n.left == nil { // 将右子树的数据反给上层
			tn := n.right
			n.left = nil
			t.size--
			return tn
		}
		// 2. 待删除节点右子树为空的情况
		if n.right == nil {
			tn := n.left
			n.left = nil
			t.size--
			return tn
		}
		// 3. 左右都有数据,
		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := t.minimum(n.right)
		//removeMin中已经 size--,外面不需要 --
		successor.right = t.removeMin(n.right)
		successor.left = n.left
		n.left = nil
		n.right = nil
		return successor
	}
}

//  返回以node为根的二分搜索树的最小值所在的节点
func (t *AVLTree) minimum(n *avlNode) *avlNode {
	if n.left == nil {
		return n
	}
	return t.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根 和 是否删除成功
func (t *AVLTree) removeMin(n *avlNode) *avlNode {
	if n.left == nil {
		// 将要删除的右节点挂载父节点上,通过返回值返给服节点
		tn := n.right
		n.right = nil // 置空 gc回收
		t.size--
		return tn
	}
	n.left = t.removeMin(n.left)
	return n
}

func (t *AVLTree) String() string {
	var sb strings.Builder
	generateBSTLevelString(t.root, &sb)
	return sb.String()
}

func generateBSTLevelString(n *avlNode, sb *strings.Builder) {
	l := list.New()
	l.PushBack(n)

	curNodeCnt := 1  // 当前行的节点
	nextNodeCnt := 0 // 下一行的节点
	for l.Len() != 0 {
		n, _ := l.Remove(l.Front()).(*avlNode)

		sb.WriteString(fmt.Sprintf("%v ", n))
		curNodeCnt--
		if n.left != nil {
			l.PushBack(n.left)
			nextNodeCnt++
		}
		if n.right != nil {
			l.PushBack(n.right)
			nextNodeCnt++
		}
		if curNodeCnt == 0 { // 当前行打印完了
			sb.WriteRune('\n')
			curNodeCnt = nextNodeCnt
			nextNodeCnt = 0
		}
	}
}
