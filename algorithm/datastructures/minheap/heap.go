// go sdk中的container包下已经有了heap堆的结构,此处实现按照java的思路实现
package minheap

// 需要使用指针接受者去实现接口
type Comparable interface {
	// 比较大小 相等返回0 , 当前这个数小返回负数 ,当前数大返回正数
	CompareTo(o Comparable) int
}

type Heap struct {
	data []Comparable //存储的数据
}

func NewHeap(d []Comparable) *Heap {
	ret := &Heap{}
	if d == nil || len(d) == 0 {
		return ret
	}
	ret.data = make([]Comparable, len(d))
	for i := 0; i < len(d); i++ {
		ret.data[i] = d[i]
	}
	//heapify 操作
	ret.heapify()
	return ret
}

func (h *Heap) heapify() {
	// 思路: 跳过叶子节点,对最小的父节点进行下沉操作,一直到根部
	// 最小的叶子节点的服节点就 parent(len()-1)
	if h.data == nil || h.Len() == 0 {
		return
	}
	for i := parent(h.Len() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap) Len() int {
	if h.data == nil {
		return 0
	}
	return len(h.data)
}

func (h *Heap) Poll() Comparable {
	//1. 取出队头元素
	//2. 将对尾元素,移到顶部
	//3. 移除尾部
	//4. 对头部下沉
	if h.Len() == 0 {
		return nil
	}
	ret := h.Peek()
	h.swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]
	h.siftDown(0)
	return ret
}

// 移除对应的元素
func (h *Heap) Remove(e Comparable, eqFunc func(e, b Comparable) bool) bool {
	// 1 找到对应元素
	var find Comparable
	var fi int
	for i := 0; i < h.Len(); i++ {
		if eqFunc(e, h.data[i]) {
			find = h.data[i]
			fi = i
			break
		}
	}
	if find == nil {
		//没有找到
		return false
	} else {
		// 与最后一个值进行替换
		h.swap(fi, h.Len()-1)
		h.data = h.data[:h.Len()-1] //移除最后一个
		// 下沉下标
		h.siftDown(fi)
		return true
	}
}

func (h *Heap) Peek() Comparable {
	if h.Len() == 0 {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Add(e Comparable) {
	h.data = append(h.data, e)
	h.siftUp(h.Len() - 1)
}

// 上浮
func (h *Heap) siftUp(i int) {
	ci := i
	pi := parent(ci)
	for ci > 0 && h.data[ci].CompareTo(h.data[pi]) < 0 {
		h.swap(pi, ci)
		ci = pi
		pi = parent(ci)
	}
}

// 下沉
func (h *Heap) siftDown(i int) {
	ci := i
	dataLen := h.Len() // 数据大小
	for leftChild(ci) < dataLen {
		// 找到较小值的孩子的下标
		mi := leftChild(ci)
		if mi+1 < dataLen && h.data[mi].CompareTo(h.data[mi+1]) > 0 { // mi + 1 表示右边下标
			// 右孩子的值小些
			mi += 1
		}
		if h.data[ci].CompareTo(h.data[mi]) <= 0 {
			// 已经比孩子小了不用下沉
			break
		}
		h.swap(mi, ci)
		ci = mi
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 返回完全二叉堆的数组表示中，一个索引所表示的元素的父亲节点的索引
func parent(index int) int {
	if index <= 0 {
		return -1 //表示没有父节点索引
	}
	return (index - 1) / 2
}

// 左孩子下标
func leftChild(index int) int {
	return index*2 + 1
}

// 右孩子下标
func rightChild(index int) int {
	return index*2 + 2
}
