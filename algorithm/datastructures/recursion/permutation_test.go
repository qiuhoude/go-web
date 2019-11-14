package recursion

import (
	"fmt"
	"strconv"
	"testing"
)

// 递归实现全排列
// 思路: 可以画出递归树图进行递归解析, abcd -> axxx, bxxx, cxxx dxxx;
//
/*
时间复杂度分析
递推公式: f(n) = n*f(n-1)
对递推公式求和: Σ = n + n*f(n-1) + n*(n-1)*f(n-2) + ...
所以的 O(n!)~O(n*n!) 之间

*/
func permutation(arr []interface{}, start, dep int) {
	len := len(arr)
	if start == len-1 { // 最后一位
		fmt.Println("dep:", strconv.Itoa(dep), arr)
	} else {
		for i := start; i < len; i++ {
			// i = start 时输出自己
			// 如果i和start的值相同就没有必要交换
			if i == start || arr[i] != arr[start] { // 减枝操作
				//交换当前这个与后面的位置
				arr[i], arr[start] = arr[start], arr[i]
				fmt.Printf("dep:%v swap1: %v, start:%v\n", dep, arr, start)
				//递归处理索引+1
				permutation(arr, start+1, dep+1)
				//换回来，因为是递归，如果不换回来会影响后面的操作，并且出现重复
				arr[i], arr[start] = arr[start], arr[i]
				fmt.Printf("dep:%v swap2: %v, start:%v\n", dep, arr, start)
			}
		}
	}
}

func Test_permutation(t *testing.T) {
	//slice1 := make([]interface{}, 4)
	//for i := 0; i < 4; i++ {
	//	slice1[i] = i + 1
	//}
	//permutation(slice1, 0)

	slice2 := make([]interface{}, 3)
	slice2[0] = "a"
	slice2[1] = "b"
	slice2[2] = "c"
	permutation(slice2, 0, 1)
	//fmt.Printf("permutation after %v\n", slice2)
}
