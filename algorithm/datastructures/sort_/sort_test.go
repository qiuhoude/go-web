package sort_

import (
	"fmt"
	"testing"
)

/*
冒泡排序
只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足
大小关系要求。如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，
重复 n 次，就完成了 n 个数据的排序工作
属于稳定性算法(相同值排序前后顺序不变)
空间 O(1)
最好 O(n)
最坏 O(n^2)

满有序度 = n*(n-1)/2
逆序度 = 满有序度 - 初始有序度
逆序度 = 就是需要要换的次数
平均逆序度 = (满有序度 - 0)/ 2  = n*(n-1)/4
平均复杂度 约等于 O(n*(n-1)/4) , n值很大时 O(n^2)

*/
func bubbleSort(arr []int) {
	n := len(arr)
	if n <= 0 {
		return
	}
	for i := 0; i < n; i++ {
		flag := false // 提前退出标志
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true // 有交换
			}
		}
		if !flag {
			break
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1}
	//bubbleSort(arr)
	fmt.Println(arr)
}

/*
插入排序

(后面排序算法思路可以借鉴)
将数组中的数据分为两个区间，已排序区间和未排序区间，初始已排序区间只有一个
元素，就是数组的第一个元素。插入算法的核心思想是取未排序区间中的元素，在已排序区间中
找到合适的插入位置将其插入，并保证已排序区间数据一直有序。重复这个过程，直到未排序区
间中元素为空，算法结束

是稳定性排序
空间 O(1)
最好 O(n)
最坏 O(n^2)
平局 O(n^2)
*/
func insertionSort(arr []int) {
	n := len(arr)
	if n < 1 {
		return
	}
	for i := 1; i < n; i++ {
		v := arr[i] // 记录要插入的值
		j := i - 1  // 0~j 表示已排好序的位置
		// 查找要插入的位置
		for ; j >= 0; j-- {
			if arr[j] > v {
				arr[j+1] = arr[j] // 移动数据
			} else {
				break
			}
		}
		arr[j+1] = v // 插入数据
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1}
	insertionSort(arr)
	fmt.Println(arr)
}

/*
选择排序

选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次
会从未排序区间中找到最小的元素，将其放到已排序区间的末尾

非稳定性的排序 , 例如 [5,8,5,2,9] ,后面的的5就会和前面的5进行换位置
空间 O(1)
最好 O(n)
最坏 O(n^2)
平局 O(n^2)
*/

func selectionSort(arr []int) {
	n := len(arr)
	if n < 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		minIndex := i            // 最小值的位置
		for j := i; j < n; j++ { // 寻找最小值位置
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1, 8, 7}
	selectionSort(arr)
	fmt.Println(arr)
}

/*
归并排序

把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了

递推公式: mergeSort(arr[p:r]) = mergeSort(arr[p:mid]) + mergeSort(arr[mid+1:r])
merge() 进行合并数组操作
mid = (p + r) / 2
终止条件 p>=r

递归代码的时间复杂度也可以写成递推公式 T(a) = T(b) + T(c) + K

T(1) = C； n=1 时，只需要常量级的执行时间，所以表示为 C。
T(n) = 2*T(n/2) + n； n>1

T(n) = 2*T(n/2) + n
	= 2*(2*T(n/4) + n/2) + n = 4*T(n/4) + 2*n
	= 4*(2*T(n/8) + n/4) + 2*n = 8*T(n/8) + 3*n
	= 8*(2*T(n/16) + n/8) + 3*n = 16*T(n/16) + 4*n
	......
	= 2^k * T(n/2^k) + k * n
	= 2*n * C 		 + log(n)*n
	= nlog(n) + 2n*C

当 T(n/2^k) = T(1) = C
n/2^k = 1  => k = log2(n)

所以任何情况下时间复杂度 O(nlog(n))

空间复杂度
2^x=n => x=log2(n)
看上去要去要分配log2(n)次,每次需要分配n个空间, O(n * log2(n));
但是同一时刻只会有n个空间被分配
所以复杂度是 O(n)
*/
func mergeSortEnter(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	mid := (p + r) / 2
	// 分而治之
	mergeSort(arr, p, mid)
	mergeSort(arr, mid+1, r)
	// 将 A[p...q] 和 A[q+1...r] 合并为 A[p...r]
	merge(arr, p, mid, r) // 归并
}

func merge(arr []int, p, mid, r int) {
	tmp := make([]int, r-p+1)
	i := p       // 前数组的指针
	j := mid + 1 // 后数组指针
	k := 0
	for ; i <= mid && j <= r; k++ {
		if arr[i] <= arr[j] { // 前数组 等于 后数组, 取前数组的值,保证稳定性
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}
	// 剩余的数据,填充到数组中
	for ; i <= mid; k++ {
		tmp[k] = arr[i]
		i++
	}
	for ; j <= r; k++ {
		tmp[k] = arr[j]
		j++
	}
	copy(arr[p:r+1], tmp)
}

func TestMergeSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1, 8, 7}
	mergeSortEnter(arr)
	t.Log(arr)
}

/*
快速排序

如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间
的任意一个数据作为 pivot（分区点）
遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放
到中间。经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间
都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的

递推公式: quickSort(arr,p,r) = quickSort(arr,p,q-1) + quickSort(arr,q+1,r)
递归出口 p>=r

空间复杂度 O(1)
最坏 O(n^2) 选取的 pivot的位置点在最边上
平均 O(nlogn)

利用的分治 和 分区的思想
*/

func quickSortEnter(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	quickSort(arr, 0, n-1)
}

func quickSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r) // 获取区服点
	quickSort(arr, p, q-1)
	quickSort(arr, q+1, r)
}

func partition(arr []int, p, r int) int {
	// 取 r位置为 pivot点
	pivot := arr[r]
	i := p // p ~ i 是<pivot
	for j := p; j < r; j++ {
		if arr[j] < pivot {
			if j != i {
				arr[j], arr[i] = arr[i], arr[j] //放到最后面
			}
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}

func TestQuickSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1, 8, 7}
	quickSortEnter(arr)
	t.Log(arr)

}

// 如何在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素
func findMaxK(arr []int, k int) int {
	n := len(arr)
	if n < 1 || n < k {
		return -1
	}
	k = n - k + 1 // 因为 partition 是左小右大的正序, 找第K大 等于 找第 n-k+1 小的元素
	r := n - 1
	p := partition(arr, 0, r)
	for j := 0; j < n; j++ { // 最多查找 n 次,防止死循环
		if p+1 == k {
			return p
		} else if k > p+1 { // 继续右边找
			p = partition(arr, p+1, r)
		} else { // 继续左边找
			p = partition(arr, 0, p-1)
		}
	}
	return -1
}

func TestFindMaxK(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1, 8, 7, 1, 1}
	k := 1
	i := findMaxK(arr, k)
	if i != -1 {
		t.Log(arr[i])
	}
}
