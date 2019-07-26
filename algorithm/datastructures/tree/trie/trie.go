// 前缀树
package trie

type Trie struct {
	root *node
	size int // 存储词的数量
}

type node struct {
	children map[string]*node // 孩子节点
	isEnd    bool             // 结束标识符
	data     interface{}      // 存储的数据
}

func NewTrie() *Trie {
	n := newNode()
	return &Trie{root: n, size: 0}
}

func newNode() *node {
	n := new(node)
	n.children = make(map[string]*node)
	return n
}

func (t Trie) Size() int {
	return t.size
}

// 添加
func (t *Trie) Add(w string, data interface{}) {
	cur := t.root
	for _, v := range []rune(w) {
		c := string(v)
		n, ok := cur.children[c]
		if !ok {
			n = newNode()
			cur.children[c] = n
		}
		cur = n
	}
	cur.data = data
	if !cur.isEnd {
		cur.isEnd = true
		t.size++
	}
}

// 删除词
func (t *Trie) Remove(w string) bool {
	cur := t.root
	var stack []*node
	runeArr := []rune(w)
	for _, v := range runeArr {
		c := string(v)
		n, ok := cur.children[c]
		if !ok { // 没有找到该单词不用删除
			return false
		}
		stack = append(stack, n)
		cur = n
	}
	// 结尾标识改掉
	cur.isEnd = false
	if len(cur.children) == 1 { // 只有自己一个字符 没有后续的 就可以进行移除操作
		for i := len(stack) - 1; i >= 0; i-- {
			c := string(runeArr[i])
			n := stack[i]
			delete(n.children, c)
			if cur.isEnd {
				break
			}
		}
	}
	return true
}

func (t *Trie) Find(w string) interface{} {
	cur := t.root
	for _, v := range []rune(w) {
		c := string(v)
		n, ok := cur.children[c]
		if !ok {
			return nil
		}
		cur = n
	}
	if !cur.isEnd { // 不是结束就返回nil
		return nil
	}
	return cur.data
}

func (t *Trie) Contains(w string) bool {
	return t.Find(w) != nil
}

// 前缀搜索
func (t *Trie) SearchPrefix(prefix string) []interface{} {
	cur := t.root
	for _, v := range []rune(prefix) {
		c := string(v)
		n, ok := cur.children[c]
		if !ok {
			return nil
		}
		cur = n
	}
	ret := make([]interface{}, 0, 8)
	searchNode(cur, &ret)
	return ret
}

func searchNode(n *node, ret *[]interface{}) {
	if n.isEnd {
		*ret = append(*ret, n.data)
	}
	if len(n.children) <= 0 {
		return
	}
	for _, cn := range n.children {
		searchNode(cn, ret)
	}
}
