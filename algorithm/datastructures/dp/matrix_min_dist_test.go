package dp

import (
	"fmt"
	"testing"
)

/*
矩阵最短距离

{
	{1,3,5,9}
	{2,1,3,4}
	{5,2,6,7}
	{6,8,4,3}
}

求解左上到右下点最短距离, 每一步只能往右 或 下走

思路: 到 (i,j) 状态的最小值,只能由(i-1,j) 或 (i,j-1) 两个状态转变过来
minDis[i][j]= m[i][j] + min(m[i][j-1], m[i-1][j])
*/

// 回溯法的写法 i表示纵轴, j 表示横轴
func minDistBacktracing(i, j, dis int, m [][]int, ret *int) {
	maxY := len(m) - 1
	maxX := len(m[0]) - 1
	if i == maxY && j == maxX {
		if *ret > dis+m[maxY][maxX] {
			*ret = dis + m[maxY][maxX]
		}
		return
	}
	if j < maxX { // 向左走
		minDistBacktracing(i, j+1, dis+m[i][j], m, ret)
	}
	if i < maxY { // 向下走
		minDistBacktracing(i+1, j, dis+m[i][j], m, ret)
	}
}

// 动态规划的方式
func minDistDp(m [][]int) int {
	yl := len(m)
	xl := len(m[0])

	dp := make([][]int, yl)
	for i := range dp {
		dp[i] = make([]int, xl)

	}
	dp[0][0] = m[0][0]
	// 设置初始首行和首列的状态
	for i := 1; i < yl; i++ { //首列
		dp[i][0] = dp[i-1][0] + m[i][0]
	}
	for j := 1; j < xl; j++ { // 首行
		dp[0][j] = dp[0][j-1] + m[0][j]
	}

	for i := 1; i < yl; i++ {
		for j := 1; j < xl; j++ {
			dp[i][j] = m[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}
	// 倒推路径
	printPath(yl-1, xl-1, m, dp)
	return dp[yl-1][xl-1]
}

func printPath(y int, x int, m [][]int, dp [][]int) {
	fmt.Printf("(%v,%v,%v)", y, x, m[y][x])
	for y > 0 && x > 0 {
		switch dp[y][x] - m[y][x] {
		case dp[y-1][x]:
			y--
		case dp[y][x-1]:
			x--
		}
		fmt.Printf("<-(%v,%v,%v)", y, x, m[y][x])
	}
	fmt.Printf("<-(%v,%v,%v)\n", 0, 0, m[0][0])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestMinDis(t *testing.T) {
	m := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
		//{6, 8, 4, 23},
	}

	ret := 0x7fffffff
	minDistBacktracing(0, 0, 0, m, &ret)
	t.Log(ret)

	ret2 := minDistDp(m)
	t.Log(ret2)

}
