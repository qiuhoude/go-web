package bm

/*
bm 字符匹配算法 Boyer-Moore
1. 好串规则 bad character rule 好串表示已匹配字符
2. 坏串规则 good suffix shift  坏串表示不匹配字符
*/

// 构建 k:字符 v:位置 的hash表
func generateBc(b []rune) map[rune]int {
	ret := make(map[rune]int, len(b))
	for index, v := range b {
		ret[v] = index
	}
	return ret
}

// 返回第一个匹配的位置
func BmSearch(mainStr, pattern []rune) int {
	if len(mainStr) == 0 || len(pattern) == 0 || len(pattern) > len(mainStr) {
		return -1
	}
	//构建坏字符哈希表
	bc := generateBc(pattern)
	n := len(mainStr)
	m := len(pattern)
	suffix, prefix := generateGS(pattern)
	step := 1
	// 主串与模式对齐的第一个字符的位置
	for i := 0; i <= n-m; i += step {

		//查找坏字符串的位置
		badCharIndex := m - 1                     // badCharIndex 表示坏字符在模式串的位置
		for ; badCharIndex >= 0; badCharIndex-- { // 从后往前进行匹配
			if mainStr[i+badCharIndex] != pattern[badCharIndex] {
				break //找到了坏字符的位置
			}
		}

		if badCharIndex <= 0 {
			// 没有坏串就匹配成功，返回主串与模式串第一个匹配的字符的位置
			return i
			/*
				// 如果此处如果要继续匹配
				step =1
				continue
			*/
		}
		/*
			1. 坏串规则:
			 i+badCharIndex是坏字符在主串的位置, mainStr[i+badCharIndex]就是坏串的字符
			 bc[mainStr[i+badCharIndex]]就是坏串字符在模式串的位置 没有就是 -1
			 badCharIndex-bc[mainStr[i+badCharIndex]] 就是要向后滑动的距离
		*/
		bcIndex := -1
		if i, ok := bc[mainStr[i+badCharIndex]]; ok {
			bcIndex = i
		}
		stepForBC := badCharIndex - bcIndex // 坏串滑动位数

		/*
			2. 好串规则
		*/
		stepForGS := -1
		if badCharIndex < m-1 { //如果有好串后缀
			stepForGS = moveByGS(badCharIndex, m, suffix, prefix) // 计算位移
		}
		// 向后移动位置
		step = max(stepForBC, stepForGS)
		if step <= 0 { // 防止负数
			step = 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

// badChartIndex 表示坏字符在模式串中的位置 ; m 表示模式串长度
func moveByGS(badCharIndex, patternLen int, suffix []int, prefix []bool) int {
	k := patternLen - 1 - badCharIndex // k 已经匹配后缀的长度
	if suffix[k] != -1 {               //complete match 直接找到位置了
		return badCharIndex - suffix[k] + 1
	}

	// 否则匹配后缀子串
	for r := badCharIndex + 2; r < patternLen-1; r++ {
		//
		if prefix[patternLen-r] { // patternLen-r 表示后缀子串的长度
			return r
		}
	}
	//no match
	return patternLen
}

// 返回好串 后缀表 与 前缀表
func generateGS(b []rune) ([]int, []bool) {
	m := len(b)
	// 存储的是该位置的字符,当前位置倒数上一次出现的位置的下标
	// cabcab 最后一个b index=5,只有1个字符 ,上一次出现的b index=2,
	// 所以 suffix[1]=2
	suffix := make([]int, m)
	// 后缀子串与前缀子串的位置是否匹配
	prefix := make([]bool, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ { //b[0;i]
		// 此处从后往前数的指针
		j := i // 从 j 位置往前数
		k := 0
		for j >= 0 && b[j] == b[m-1-k] { // 与 b[0, m-1] 求公共后缀子串
			// m-1-k 理解为从末尾往前数
			j--
			k++
			// k 表示长度
			suffix[k] = j + 1 //j+1 表示公共后缀子串在 b[0, i] 中的起始下标
		}
		if j == -1 {
			prefix[k] = true // 如果公共后缀子串也是模式串的前缀子串
		}
	}
	return suffix, prefix
}
