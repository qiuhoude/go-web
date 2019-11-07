package minheap

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Task struct {
	Name     string
	Priority int
}

func taskCompare(v1, v2 interface{}) int {
	t1, _ := v1.(*Task)
	t2, _ := v2.(*Task)
	return t1.Priority - t2.Priority
}

func eqTaskFunc(e, b interface{}) bool {
	t1, ok1 := e.(*Task)
	t2, ok2 := b.(*Task)
	if ok1 && ok2 && t1 == t2 {
		fmt.Println("删除了 ", t1.Name)
		return true
	}
	return false
}

func generatorTask() []interface{} {
	rand.Seed(time.Now().Unix())
	const size = 10
	d := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		p := rand.Intn(300)
		d = append(d, &Task{
			Name:     "task-" + strconv.Itoa(i),
			Priority: p,
		})
	}
	return d
}

func TestHeap_Remove(t *testing.T) {
	tasks := generatorTask()
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
	fmt.Println("----------------")
	h := NewHeap(taskCompare)
	h.Heapify(tasks...)
	remove := h.Remove(tasks[1], eqTaskFunc)

	fmt.Println("删除 ", remove)
	for h.Len() != 0 {
		e := h.Poll()
		if t, ok := e.(*Task); ok {
			fmt.Println(t)
		}
	}
}

func TestHeap_Add(t *testing.T) {
	h := NewHeap(taskCompare)
	h.Heapify(generatorTask()...)
	cnt := 0
	for h.Len() != 0 {
		cnt++
		pe := h.Peek().(*Task)
		pe.Name = pe.Name + "_p" + strconv.Itoa(cnt)
		e := h.Poll()
		if t, ok := e.(*Task); ok {
			fmt.Println(t)
		}
	}
}

func Test_topKFrequent(t *testing.T) {
	ret := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	t.Log(ret)
}

// 通过heap 求中位数 , 百分位等问题
// 思路: 维护两个堆, 一个最大堆(保存前半部数据)，一个最小堆(保存后半部分)
// 动态维护: 1 比较插入值的大小,小于大堆top进大堆, 大于小堆top进小堆, 在大堆小堆之间得范围随便插入到哪
// 		   2 维护两堆的数量平衡, 通过堆之间的数据搬迁，使得两个堆的size比例满足中位数 或 百分位的比例

// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {

	// 统计
	freqMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := freqMap[v]; ok {
			freqMap[v]++
		} else {
			freqMap[v] = 1
		}
	}
	var fdd []numFrequent
	for k, v := range freqMap {
		fdd = append(fdd, numFrequent{num: k, freq: v})
	}

	h := NewHeap(func(v1, v2 interface{}) int {
		i1 := v1.(numFrequent)
		i2 := v2.(numFrequent)
		if i1.freq > i2.freq {
			return 1
		} else if i1.freq < i2.freq {
			return -1
		} else {
			return 0
		}
	})
	for _, v := range fdd {
		if h.Len() < k {
			h.Add(v)
		} else if h.Peek().(numFrequent).freq < v.freq {
			h.Poll()
			h.Add(v)
		}
	}
	var ret []int
	for h.Peek() != nil {
		ret = append(ret, h.Poll().(numFrequent).num)
	}
	return ret
}

type numFrequent struct {
	num  int
	freq int
}
