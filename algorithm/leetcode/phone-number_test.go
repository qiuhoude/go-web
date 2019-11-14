package leetcode

import "testing"

func TestLetterCombinations(t *testing.T) {
	res := letterCombinations("213")
	t.Log(res)
}

// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letterMap := make(map[rune][]rune)
	letterMap['2'] = []rune{'a', 'b', 'c'}
	letterMap['3'] = []rune{'d', 'e', 'f'}
	letterMap['4'] = []rune{'g', 'h', 'i'}
	letterMap['5'] = []rune{'j', 'k', 'l'}
	letterMap['6'] = []rune{'m', 'n', 'o'}
	letterMap['7'] = []rune{'p', 'q', 'r', 's'}
	letterMap['8'] = []rune{'t', 'u', 'v'}
	letterMap['9'] = []rune{'w', 'x', 'y', 'z'}
	var s []rune
	var res []string
	findCombination(0, []rune(digits), letterMap, s, &res)
	return res
}

func findCombination(index int, text []rune, letterMap map[rune][]rune, s []rune, res *[]string) {
	if index == len(text) {
		// 递归的出口就是解
		*res = append(*res, string(s))
		return
	}
	arr, ok := letterMap[text[index]]
	if !ok {
		// 没有找到就直接继续
		findCombination(index+1, text, letterMap, s, res)
		return
	}
	for i := range arr {
		findCombination(index+1, text, letterMap, append(s, arr[i]), res)
	}
}
