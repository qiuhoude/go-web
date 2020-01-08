package rsa

// 求两个数的最大公约数
// 欧几里 辗转相除法
func greatestCommonDivisor1(a, b int) int {
	// 欧几里的辗转相除法
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 九章算术 更相减损术
func greatestCommonDivisor2(a, b int) int {
	// 两个正整数a和b（a>b），它们的最大公约数等于a-b的差值c和较小数b的最大公约数。
	// 比如10和25，25减去10的差是15,那么10和25的最大公约数，等同于10和15的最大公约数。
	if a < b {
		a, b = b, a
	}
	for a != b {
		a = a - b
		if a < b {
			a, b = b, a
		}
		a, b = b, a-b
	}
	return b
}

func greatestCommonDivisor3(a, b int) int {
	// 更相减损术与移位结合
	if a == b {
		return b
	}
	if a < b { // 保证a >= b
		a, b = b, a
	}

	if a&1 != 0 && b&1 != 0 {
		//当a和b均为奇数，利用更相减损术运算一次，gcb(a,b) = gcb(b, a-b)， 此时a-b必然是偶数，又可以继续进行移位运算
		return greatestCommonDivisor3(b, a-b)
	} else if a&1 == 0 && b&1 == 0 {
		// 当a和b均为偶数，gcb(a,b) = 2*gcb(a/2, b/2) = 2*gcb(a>>1, b>>1)
		return 2 * greatestCommonDivisor3(a>>1, b>>1)
	} else if a&1 == 0 && b&1 != 0 {
		// 当a为偶数，b为奇数，gcb(a,b) = gcb(a/2, b) = gcb(a>>1, b)
		return greatestCommonDivisor3(a>>1, b)
	} else {
		// 当a为奇数，b为偶数，gcb(a,b) = gcb(a, b/2) = gcb(a, b>>1)
		return greatestCommonDivisor3(a, b>>1)
	}
}
