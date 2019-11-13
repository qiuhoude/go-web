package kmp

/*

kmp算法步骤
1.  生成前缀表 (此处kmp算法是根据阮一峰讲解的kmp算法实现的,可能会与其他kmp算法的求前缀表有所区别)
	阮一峰的方式 是指针指向主串,主串指针可以移动一大段,每次比较都是从子串第一个开始 与 主串指针的位置进行比较
	其他人讲解的 是指针指向模式串, 通过变动模式串的指针,每次从模式串指针的位置开始(不必每次从字串第一个开始),与 主串指针(每次for都只+1)进行比较
2.  统计已匹配字符数
3.  查表,根据公式往下一点:移动位数 = 已匹配的字符数 - 对应的部分匹配值(查表获得)


时间复杂度 O(m+n) m是主串的长度, n是匹配串的长度
*/
// 匹配主中有多少个子串
func MarchSubstr(mainStr, pattern []rune) int {
	if len(mainStr) < len(pattern) { // 子串大于主肯定没有匹配的
		return 0
	}
	preTab := prefixTable2(pattern) //前缀表
	var (
		mp       = 0 // 主串的下标
		mLen     = len(mainStr)
		sLen     = len(pattern)
		marchCnt = 0 // 已匹配数量
	)
	for mp+sLen <= mLen {
		mCnt := 0 //已匹配的字符个数
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
	table[0] = 0
	for i := 1; i < sLen; i++ {
		// 表示 [0,i]的子串 中最长公共前后缀的的长度
		table[i] = findMaxCommonPreSufNum(s[:i+1])
		//fmt.Printf("%s=%d\n", string(s[:i+1]), table[i])
	}
	return table
}

func prefixTable2(s []rune) []int {
	m := len(s)
	nexts := make([]int, m)
	for index := range nexts {
		nexts[index] = -1
	}

	for i := 1; i < m-1; i++ {
		j := nexts[i-1] // 上一次匹配前缀与后缀匹配最大长度,前缀末尾位置
		// j+1 前缀的末尾位置的后面一位  比如 ababacd 模式串, 这时 i=4时s[i]='a',j=1, j+1的位置就是第2个a
		for s[j+1] != s[i] && j >= 0 {
			// 如果 目s[j+1] 表示最大前缀后一位字符 不等于 当前字符,就进去查找次匹配长度的位置(相当于查找了已经求解过的子问)
			// 看次长度的后面一位 是否相等 当前字符
			j = nexts[j]
		}

		if s[j+1] == s[i] {
			j += 1
		}

		nexts[i] = j
	}
	// 将下标 改为长度
	table := make([]int, m)
	for i := range nexts {
		table[i] = 1 + nexts[i]
	}
	//fmt.Println(table)
	return table
}

// 查找前缀与后缀的共有最大长度
func findMaxCommonPreSufNum(s []rune) int {
	// 此处的思路,每次从最长的进行比较,
	// 可以通过动态规划, 从最短的开始比较进行递推
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
