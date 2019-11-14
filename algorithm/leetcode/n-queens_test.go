package leetcode

import (
	"fmt"
	"testing"
)

// 51. N皇后
// https://leetcode-cn.com/problems/n-queens/

func TestSolveNQueens(t *testing.T) {
	t.Log(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	ret := make([]int, n)
	var retStr [][]string
	calQueens(0, n, ret, &retStr)
	return retStr
}

func calQueens(row, n int, ret []int, retStr *[][]string) {
	if row >= n {
		// 找到解
		solve := make([]string, n)
		c := make([]rune, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if ret[i] == j {
					c[j] = 'Q'
				} else {
					c[j] = '.'
				}
			}
			solve[i] = string(c)
			fmt.Println(string(c))
		}
		*retStr = append(*retStr, solve)
		fmt.Println("-----------")
		return
	}
	for col := 0; col < n; col++ { // 每行都有n种可能
		if isOk(row, col, n, ret) {
			ret[row] = col                   // 第 row 行的棋子放到了 column 列
			calQueens(row+1, n, ret, retStr) // 如果满足条件就进行尝试下一行
		}
	}
}

// 检测横竖是否有其他皇后
func isOk(row, col, n int, ret []int) bool {
	lCol := col - 1 // 左边列
	rCol := col + 1 // 右边列

	for r := row - 1; r >= 0; r-- {
		if ret[r] == col { // 纵向有相同的
			return false
		}
		// 左上对角线
		if lCol >= 0 && lCol == ret[r] {
			return false
		}
		// 右上对角线
		if rCol < n && rCol == ret[r] {
			return false
		}
		lCol--
		rCol++
	}
	return true
}
