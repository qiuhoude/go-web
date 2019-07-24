package leetcode

import (
	"fmt"
	"testing"
)

//上台阶问题
//n 阶楼梯 每次上1台阶或2台阶,共有多少种上法?

func TestClimbStairs(t *testing.T) {
	n := 4
	fmt.Println(calcWays(n))
	fmt.Println(calcWays2(n))
}

// 递归的写法
func calcWays(n int) int {
	if n == 1 { // 最后一个是1个台阶
		return 1
	}
	if n == 2 {
		return 2
	}
	return calcWays(n-1) + calcWays(n-2)

}

func calcWays2(n int) int {
	// 记忆化搜索的方式 自顶向下的思维方式
	if n == 1 { // 最后一个是1个台阶
		return 1
	}
	v := make([]int, n+1)
	v[0] = 1
	v[1] = 1
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			v[i] = v[i-1] + v[i-2]
		}
	}
	return v[n]
}

func TestFib(t *testing.T) {
	n := 5
	t.Log(fib(n))
	t.Log(fib2(n))
	t.Log(fib3(n))
}

// 递归写法
func fib2(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	// 此处fib2 会有很多重复计算
	return fib2(n-1) + fib2(n-2)
}

// 动态规划的方式
func fib3(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	// 数组记录前n
	v := make([]int, n+1)
	v[1] = 1
	v[2] = 1
	for i := 3; i <= n; i++ {
		if v[i] == 0 {
			v[i] = v[i-1] + v[i-2]
		}
	}

	return v[n]
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	n1 := 1
	n2 := 1
	for i := 0; i < n-1; i++ {
		n1, n2 = n2, n1+n2
	}
	return n1
}
