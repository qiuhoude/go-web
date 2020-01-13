package binary_search

import "math"

/*
二分查找
有点类似于数学中的夹逼定理, 两边不断逼近某个值

十个二分九个错

时间复杂度
2^k = n 第 k次可以找到该元素
k = log2n 所以 O(logn)

*/
func BSearch(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high { // 注意是 <= 不是 <
		mid := low + ((high - low) >> 1)
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			low = mid + 1 // 注意需要+1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 二分查找递归写法
func BSearchRecusion(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	return bsearchR(arr, val, 0, n-1)
}

func bsearchR(arr []int, val, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if arr[mid] == val {
		return mid
	} else if arr[mid] < val {
		return bsearchR(arr, val, mid+1, high)
	} else {
		return bsearchR(arr, val, low, mid-1)
	}
}

// 二分法求平方根
func BSqrt(x float64, precise float64) float64 {
	if x <= 0 {
		if x == 0 {
			return 0
		}
		return -1
	}
	// 大于1 在 0~x 之间查找
	low := 0.0
	high := x
	if x < 1 {
		// 小于1 在 x~1 之间查找该值
		low = x
		high = 1
	}
	mid := x / 2.0
	for math.Abs(mid*mid-x) > precise {
		if mid*mid < x {
			low = mid
		} else {
			high = mid
		}
		mid = (low + high) / 2.0
	}
	return mid
}

/*
二分查找的 变体(里面有重复值时)
变体1. 查找第一个值等于给定值的元素
*/
func BSearchFirst(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			//如果 mid 等于 0，那这个元素已经是数组的第一个元素，那它肯定是我们要找的
			//如果 mid 不等于 0，但 a[mid] 的前一个元素 a[mid-1] 不等于 value，那也说明
			//a[mid] 就是我们要找的第一个值等于给定值的元素
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				// 需要找的范围肯定 [low,mid]范围,但需要找第一个所以往前逼近
				// 所以近一步向前缩小范围 [low, mid -1]
				high = mid - 1
			}
		}
	}
	return -1
}

// 烧脑的写法
func BSearchFirst2(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if arr[low] == val {
		return low
	}
	return -1
}

/*
变体2：查找最后一个值等于给定值的元素
*/
func BSearchLast(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			// 数组最后一个, 后面一个又不等于需要的值
			if mid == n-1 || arr[mid+1] != val {
				return mid
			} else { //向后逼近
				low = mid + 1
			}
		}
	}
	return -1
}

/*
变体3：查找第一个大于等于给定值的元素
*/
func BSearchFirstGeVal(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= val { // 大于等于某个值
			// 数组的第一个 ,前面一个又小于该值
			if mid == 0 || arr[mid-1] < val {
				return mid
			} else {
				// 向前逼近
				high = mid - 1
			}
		} else if arr[mid] < val {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个大于 给定值的下标
func BSearchFirstGVal(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

/*
变体4：查找最后一个小于等于给定值的元素
*/
func BSearchLastLeVal(arr []int, val int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] <= val {
			// 数组最后一个, 后面一个又大于需要的值
			if mid == n-1 || arr[mid+1] > val {
				return mid
			} else { //向后逼近
				low = mid + 1
			}
		}
	}
	return -1
}
