package sort_

import "testing"

/*
线性级别的排序
但对数据是有要求的
*/

/*
桶排序（Bucket sort）
桶排序比较适合用在外部排序中，所谓的外部排序就是数据存储在外部磁盘中，数据量比较大，
内存有限，无法将数据全部加载到内存中
此思路可以用于多进程或多线程 对一个大数组进行排序, 核心就是进行数据切分

步骤
1. 找出全局最大值
2. 进行分桶,将数据放到各自区域的桶中
3. 桶内进行排序
*/

// 获取待排序数组中的最大值
func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func bucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max := getMax(a)
	buckets := make([][]int, num) // 二维切片

	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max                // 桶序号
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			quickSortEnter(buckets[i]) // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

func TestBucketSort(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4}
	bucketSort(a)
	t.Log(a)
}

/*
计数排序（Counting sort）

计数排序只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大
很多，就不适合用计数排序了。而且，计数排序只能给非负整数排序，如果要排序的数据是其他
类型的，要将其在不改变相对大小的情况下，转化为非负整数
*/

// 假设数组中存储的都是非负整数
func countingSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	max := getMax(arr)
	c := make([]int, max+1)

	// 计算每个元素的个数，放入 c 中
	for i := 0; i < n; i++ {
		c[arr[i]]++
	}

	// 累加
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}
	// 临时数组 r，存储排序之后的结果
	tmp := make([]int, n)

	// 核心部分
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		tmp[index] = arr[i]
		c[arr[i]]--
	}

	copy(arr, tmp)
}

func TestCountingSort(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4}
	countingSort(a)
	t.Log(a)
}

/*
基数排序（Radix sort）
分割维度进排序,维度之间是递进关系
比如 手机号排序,先按照最后一位来排序手
机号码，然后，再按照倒数第二位重新排序，以此类推，最后按照第一位重新排序。经过 11 次
排序之后，手机号码就都有序了
*/
