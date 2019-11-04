package sikplist

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

/*
跳表
时间复杂度

1. 假设 每2个节点一个索引,最顶层索引是2, 有n个节点,每层最多遍历2+1=3个
设 n 个节点 , h表示层数, m表示每层遍历的次数
n/(2^h) = 2 => h = log2(n)-1
所以复杂度是 O(m*(log2(n)-1)) m是个常数级别
时间复杂度是 O(logn)

每3个节点一个索引  O(m*(log3(n)-1)); m = 3+1

空间复杂度
2节点一索引: n/2 + n/4 + n/8 +... 等比数列 约等于  约等于 n
所以空间复杂度 O(n)

3节点一索引: n/3 + n/9 + n/27 +... 等比数列 约等于 约等于 n/2

*/
const (
	MaxLevel  = 16
	SkiplistP = 0.5
)

//跳表节点结构体
type skipListNode struct {
	v        interface{}     // 数据
	score    int             // 用于排序的分值
	level    int             //层高
	forwards []*skipListNode //每层前进指针
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{
		v:        v,
		score:    score,
		forwards: make([]*skipListNode, level, level),
		level:    level,
	}
}

func (sln *skipListNode) String() string {
	return fmt.Sprintf("[%v %v]", sln.v, sln.score)
}

// 跳表结构体
type SkipList struct {
	head   *skipListNode //跳表头结点
	level  int           //跳表当前层数
	length int           ///跳表长度
}

func NewSkipList() *SkipList {
	// 头结点, 可以称为哨兵
	head := newSkipListNode(nil, math.MinInt32, MaxLevel)
	return &SkipList{head, 1, 0}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

//插入节点到跳表中
func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}
	//查找插入位置
	cur := sl.head
	//创建数组 记录每层的路径
	update := [MaxLevel]*skipListNode{}
	// 从最高层开始
	for i := MaxLevel - 1; i >= 0; i-- {
		for ; nil != cur.forwards[i]; cur = cur.forwards[i] { // 当前层进行查找
			if cur.forwards[i].v == v || cur.v == v { // 插入的值相等
				return 2
			}
			if cur.forwards[i].score > score { // 找到了
				update[i] = cur
				break
			}
		}
		update[i] = cur
	}

	//通过随机算法获取该节点层数
	level := sl.randomLevel()

	//创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	//给每层插入新结点
	for i := 0; i < level; i++ {
		newNode.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = newNode
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++
	return 0
}

/*
一级索引中元素个数应该占原始数据的 1/2，二级索引中元素个数占 1/4，三级索引 1/8 一直到最顶层。
 因为这里每一层的晋升概率是 50%。对于每一个新插入的节点，都需要调用 randomLevel 生成一个合理的层数。
该 randomLevel 方法会随机生成 1~MAX_LEVEL 之间的数，
	50%的概率返回 1
  	25%的概率返回 2
  	12.5%的概率返回 3 ...
*/
func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Float32() < SkiplistP && level < MaxLevel {
		level++
	}
	return level
}

//查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for ; nil != cur.forwards[i] && cur.forwards[i].score <= score; cur = cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			}
		}
	}
	return nil
}

// 范围查找
func (sl *SkipList) FindRange(start, end int) []interface{} {
	if sl.length == 0 || end < start {
		return nil
	}
	cur := sl.head.forwards[sl.level-1]
	for i := sl.level - 1; i >= 0; i-- {
		for ; nil != cur; cur = cur.forwards[i] {
			if cur.score >= start {
				break
			}
		}
	}
	if cur != nil {
		var ret []interface{}
		for ; cur != nil && cur.score <= end; cur = cur.forwards[0] {
			ret = append(ret, cur.v)
		}
		return ret
	}
	return nil
}

//删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}
	cur := sl.head
	//记录前驱路径
	update := [MaxLevel]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head // 默认是头结点开始
		for ; nil != cur.forwards[i]; cur = cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil { //当前层没有元素
			sl.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--
	return 0
}

func (sl *SkipList) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("level:%+v, length:%+v\n", sl.level, sl.length))
	for i := MaxLevel - 1; i >= 0; i-- {
		cur := sl.head.forwards[i]
		if cur == nil {
			continue
		}
		_, _ = fmt.Fprintf(&sb, "%v [head]->", i)
		for ; cur != nil; cur = cur.forwards[i] {
			_, _ = fmt.Fprintf(&sb, "[%v %v]->", cur.v, cur.score)
		}
		sb.WriteString("nil\n")
	}
	return sb.String()
}
