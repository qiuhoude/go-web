package dp

import (
	"fmt"
	"strconv"
	"testing"
)

/*
字符校正


中文字符的校正
中文纠错很多时候是通过拼音进行的，比如 "刘得花"->"liudehua"->"刘德华".
拼音检索方法也有很多，比如可以把热门词汇的拼音字母组织成Trie树，每个热词的结尾汉字
的最后一个拼音字母就是叶子，整体性能就是O(n)的，n为query的拼音总长度. 除了拼音外
也有根据字形（二维文字版的编辑距离？）甚至语义等做的纠错策略。
传统搜索引擎中的查询词智能提示、纠错、同义词、近义词、同好词、相关搜索、知识图谱
等系列功能统称为用户的意图识别模块。

*/

// 1. 莱温斯坦距离 Levenshtein distance

/*
增加一个字符、删除一个字符、替换一个字符 编辑距离都+1 也就是 str1[i]!=str2[i] 编辑距离就+1
回溯法方式
i 串i的指针
j 串j的指针
editDis 编辑距离
minDis 最小编辑距离
*/
func lvstBT(i, j, editDis int, stri, strj []rune, minDis *int) {
	leni := len(stri)
	lenj := len(strj)
	if i == leni || j == lenj { // 出口
		if i < leni {
			editDis += leni - i
		}
		if j < lenj {
			editDis += lenj - j
		}
		if editDis < *minDis {
			*minDis = editDis
		}
		return
	}
	if stri[i] == strj[j] { // 相等编辑距离 +0
		lvstBT(i+1, j+1, editDis, stri, strj, minDis)
	} else { // 不相等的情况下 , 就进 (i+1,j) (i,j+1) (i+1,j+1) 比较尝试
		lvstBT(i+1, j, editDis+1, stri, strj, minDis)
		lvstBT(i, j+1, editDis+1, stri, strj, minDis)
		lvstBT(i+1, j+1, editDis+1, stri, strj, minDis)
	}
}

// 动态规划的写法 ,
// 递推式  stri[i]==strj[j] minDis = min(dis(i-1,j) , dis(i,j-1) ,dis(i-1,j-1))
// 		  stri[i]!=strj[j] minDis = min(dis(i-1,j) , dis(i,j-1) ,dis(i-1,j-1))+1
func lvstDP(stri, strj []rune) int {
	leni := len(stri)
	lenj := len(strj)
	dp := make([][]int, leni)
	for j := range dp {
		dp[j] = make([]int, lenj)
	}
	if stri[0] == strj[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	// 求出首行 首列的值
	for i := 1; i < leni; i++ { //首列
		if stri[i] == strj[0] {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}
	for j := 1; j < lenj; j++ { //首行
		if stri[0] == strj[j] {
			dp[0][j] = dp[0][j-1]
		} else {
			dp[0][j] = dp[0][j-1] + 1
		}
	}
	// 依次填表
	for i := 1; i < leni; i++ {
		for j := 1; j < lenj; j++ {
			if stri[i] == strj[j] {
				dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[leni-1][lenj-1]
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}

func Test_lvst(t *testing.T) {
	stri := []rune("mitcmu")
	strj := []rune("mtacnu")
	editDis1 := strconv.IntSize
	lvstBT(0, 0, 0, stri, strj, &editDis1)
	t.Log("最小编辑距离 editDis1:", editDis1)

	editDis2 := lvstDP(stri, strj)
	t.Log("最小编辑距离 editDis2:", editDis2)
}

// 2. 最长公共子串 Longest common substring length
// 最长公共子串长度只允许增加、删除字符这两个编辑操作
// 只要当 str1[i] == str2[i] 或str1[i] == str2[i+1] 或str1[i+1] == str2[i] 这三种情况都算长度+1
// 回溯写法
func lcsBT(i, j, lcs int, stri, strj []rune, maxLen *int) {
	leni := len(stri)
	lenj := len(strj)
	if i == leni || j == lenj { // 出口
		if lcs > *maxLen {
			*maxLen = lcs
		}
		return
	}
	if stri[i] == strj[j] {
		lcsBT(i+1, j+1, lcs+1, stri, strj, maxLen)
	} else {
		lcsBT(i, j+1, lcs, stri, strj, maxLen)
		lcsBT(i+1, j, lcs, stri, strj, maxLen)
		lcsBT(i+1, j+1, lcs, stri, strj, maxLen)
	}
}

// 动态规划的方式
func lcsDP(stri, strj []rune) int {
	leni := len(stri)
	lenj := len(strj)
	dp := make([][]int, leni)
	for j := range dp {
		dp[j] = make([]int, lenj)
	}
	if stri[0] == strj[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	// 求出首行 首列的值
	for i := 1; i < leni; i++ { //首列
		if stri[i] == strj[0] && i == 1 {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < lenj; j++ { //首行
		if stri[0] == strj[j] && j == 1 {
			dp[0][j] = dp[0][j-1] + 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < leni; i++ {
		for j := 1; j < lenj; j++ {
			if stri[i] == strj[j] {
				dp[i][j] = max3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			} else {
				dp[i][j] = max3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[leni-1][lenj-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(a, max(b, c))
}

func Test_lcs(t *testing.T) {
	stri := []rune("mitcmu")
	strj := []rune("mtacnu")
	lsc1 := 0
	lcsBT(0, 0, 0, stri, strj, &lsc1)
	t.Log(" 最长公共子串 lsc1:", lsc1)

	lsc2 := lcsDP(stri, strj)
	t.Log(" 最长公共子串 lsc2:", lsc2)
}

/*
我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？比如
2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列
的长度是 4
*/

func longestSubsequence(arr []int) int {
	n := len(arr)
	dp := make([]int, n) // dp 下标的状态当数组长度, 里面的值表示数组当前长度是对于最大的递增串的长度
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i] >= arr[j] { //倒序遍历,只要找到 后面比前面大的 就+1
				dp[i] = dp[j] + 1
				break
			} else if dp[i] == 1 {
				dp[i] = dp[i-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}

func Test_longestSubsequenceBT(t *testing.T) {
	arr := []int{2, 9, 3, 6, 5, 1, 7}
	ret := longestSubsequence(arr)
	t.Log("ret:", ret)
}
