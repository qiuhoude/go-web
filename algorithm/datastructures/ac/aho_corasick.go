package ac

import (
	"github.com/qiuhoude/go-web/algorithm/datastructures/queue"
)

// 节点
type acNode struct {
	children map[rune]*acNode // 孩子节点
	isEnd    bool             // 结束标识符
	fail     *acNode          // 失败节点
	length   int
}

func newAcNode() *acNode {
	return &acNode{
		children: make(map[rune]*acNode),
	}
}

// Ac自动机 Aho-Corasick
type AC struct {
	root *acNode // 根节点
	size int     // 自动机中词的数量
}

func NewAc() *AC {
	return &AC{
		root: newAcNode(),
		size: 0,
	}
}

// 添加词
func (ac *AC) AddWorld(w string) {
	cur := ac.root
	aw := []rune(w)
	for _, v := range aw {
		n, ok := cur.children[v]
		if !ok {
			n = newAcNode()
			cur.children[v] = n
		}
		cur = n
	}
	if !cur.isEnd {
		cur.isEnd = true
		cur.length = len(aw) //长度
		ac.size++
	}
}

func (ac *AC) AddWorlds(worlds []string) {
	for _, v := range worlds {
		ac.AddWorld(v)
	}
}

// 构建失败节点
func (ac *AC) BuildFailurePointer() {
	que := queue.NewLinkedQueue()
	ac.root.fail = nil // 根节点的失败节点是nil
	que.Enqueue(ac.root)
	for !que.IsEmpty() {
		cur := que.Dequeue().(*acNode)
		if cur == ac.root { // 取出来的是root,就将子节点的的fail节点指向root
			for _, v := range cur.children {
				v.fail = cur
				que.Enqueue(v)
			}
		} else {
			for k, v := range cur.children {
				failTo := cur.fail
				for failTo != nil {
					if node, ok := failTo.children[k]; ok {
						v.fail = node
						break
					} else {
						failTo = failTo.fail
					}
				}
				if failTo == nil {
					v.fail = ac.root
				}
				que.Enqueue(v)
			}
		}
	}
}

// 匹配
func (ac *AC) Match(s string, f func(start, end int)) {
	text := []rune(s)
	textLen := len(text)
	p := ac.root
	for i := 0; i < textLen; i++ {
		c := text[i]
		var ok bool
		for _, ok = p.children[c]; !ok && p != ac.root; {
			// 不匹配, 一直往前找,直到匹配或到root节点
			p = p.fail
			_, ok = p.children[c]
		}
		if p, ok = p.children[c]; !ok {
			p = ac.root
			continue
		}
		cur := p
		for cur != ac.root {
			// 此处加个for 属于匹配更多的屏蔽词
			if cur.isEnd {
				// 找到了匹配的位置
				f(i-cur.length+1, i)
			}
			cur = cur.fail
		}
	}
}

// 删除词
func (ac *AC) Remove(w string) bool {
	cur := ac.root
	var stack []*acNode
	runeArr := []rune(w)
	for _, v := range runeArr {

		n, ok := cur.children[v]
		if !ok { // 没有找到该单词不用删除
			return false
		}
		stack = append(stack, n)
		cur = n
	}
	ac.size--
	// 结尾标识改掉
	cur.isEnd = false
	if len(cur.children) == 1 { // 只有自己一个字符 没有后续的 就可以进行移除操作
		for i := len(stack) - 1; i >= 0; i-- {
			c := runeArr[i]
			n := stack[i]
			delete(n.children, c)
			if cur.isEnd {
				break
			}
		}
	}
	return true
}

func (ac *AC) Contains(w string) bool {
	cur := ac.root
	for _, v := range []rune(w) {

		n, ok := cur.children[v]
		if !ok {
			return false
		}
		cur = n
	}
	if !cur.isEnd { // 不是结束就返回nil
		return false
	}
	return true
}
