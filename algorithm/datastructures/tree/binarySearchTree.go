package tree

import (
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

//前序遍历非递归的方式 NR=non recursion
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

// 非递归的中序遍历
func (t *BST) InOrderNR(f TraverseFunc) {
	//if t.root == nil {
	//	return
	//}
	//
	//s := stack.New()
	//s.Push(t.root)
	//
	//for !s.IsEmpty() {
	//	n, _ := s.Peek().(*node)
	//	if n.left != nil {
	//		s.Push(n.left)
	//		continue
	//	}
	//	tn, _ := s.Pop().(*node)
	//	f(tn.val)
	//	if tn.right != nil {
	//		s.Push(tn.right)
	//	}
	//}
}

func (t *BST) String() string {
	var sb strings.Builder
	generateBSTString(t.root, 0, &sb)
	return sb.String()
}

func generateBSTString(n *node, depth int, sb *strings.Builder) {
	if n == nil {
		generateDepthString(depth, sb)
		sb.WriteString("nil\n")
		return
	}
	generateDepthString(depth, sb)
	if n.val != nil {
		sb.WriteString(fmt.Sprintf("val:%v,size:%d,depth%d\n", n.val, n.size, n.depth))
	}
	generateBSTString(n.left, depth+1, sb)
	generateBSTString(n.right, depth+1, sb)
}

func generateDepthString(depth int, sb *strings.Builder) {
	for i := 0; i < depth; i++ {
		sb.WriteString("--")
	}
}
