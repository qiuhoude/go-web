package kmp

/*
kmp算法步骤
1.
2.

*/
// 匹配主中有多少个子串
func MarchSubstr(mainStr, pattern []rune) int {
	if len(mainStr) < len(pattern) { // 子串大于主肯定没有匹配的
		return 0
	}
	preTab := prefixTable(pattern) //前缀表
	var (
		mp       = 0 // 主串的下标
		mLen     = len(mainStr)
		sLen     = len(pattern)
		marchCnt = 0 // 已匹配数量
	)
	for mp+sLen <= mLen {
		mCnt := 0 //已匹配的字符数
		for sp := 0; sp < sLen; sp++ {
			if pattern[sp] != mainStr[mp+sp] {
				break
			}
			mCnt++
		}
		if mCnt == sLen { //说明匹配成功
			//fmt.Printf("匹配成功位置: %d, str:%s \n", mp, string(mainStr[mp:mp+sLen]))
			marchCnt++
			mp++
		} else if mCnt == 0 { // 没有匹配往后移一格
			mp++
		} else {
			// 移动位数 = 已匹配的字符数 - 对应的部分匹配值
			mp += mCnt - preTab[mCnt-1]
		}
	}
	return marchCnt
}

// 生成前缀表 PrefixTable
func prefixTable(s []rune) []int {
	sLen := len(s)
	table := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		table[i] = findMaxCommonPreSufNum(s[:i+1])
		//fmt.Printf("%s=%d\n", string(s[:i+1]), table[i])
	}
	return table
}

// 查找前缀与后缀的共有最大长度
func findMaxCommonPreSufNum(s []rune) int {
	length := len(s)
	maxLen := length - 1
	for maxLen > 0 {
		if eqRuneSlice(s[:maxLen], s[length-maxLen:]) {
			return maxLen
		}
		maxLen--
	}
	return 0
}

// 比较rune切片元素相等情况
func eqRuneSlice(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
