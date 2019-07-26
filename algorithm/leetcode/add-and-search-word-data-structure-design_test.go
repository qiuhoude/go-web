package leetcode

import (
	"testing"
)

// 211. 添加与搜索单词 - 数据结构设计
// https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/

type WordDictionary struct {
	Next  map[rune]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	next := make(map[rune]*WordDictionary)
	return WordDictionary{Next: next, isEnd: false}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	rArr := []rune(word)
	cur := this
	for _, c := range rArr {
		next, ok := cur.Next[c]
		if !ok {
			n := Constructor()
			next = &n
			cur.Next[c] = next
		}
		cur = next
	}
	cur.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	return march(this, []rune(word), 0)
}

func march(n *WordDictionary, word []rune, index int) bool {
	if len(word) == index {
		return n.isEnd
	}
	c := word[index]
	if c != '.' {
		next, ok := n.Next[c]
		if !ok {
			return false
		} else {
			return march(next, []rune(word), index+1)
		}
	} else {
		for _, next := range n.Next { //匹配所有边
			if march(next, []rune(word), index+1) { // 只要有一个匹配就返回
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func TestSearch(t *testing.T) {
	type args struct {
		word    []string
		searchW []string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"xx", args{[]string{"bad", "dad", "mad"},
			[]string{"", "pad", "bad", ".ad", "b.."}},
			[]bool{false, false, true, true, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			for _, w := range tt.args.word {
				obj.AddWord(w)
			}
			for i, v := range tt.args.searchW {
				got := obj.Search(v)
				if tt.want[i] != got {
					t.Errorf("%s search %d %s got %v want %v", tt.name, i, v, got, tt.want[i])
				}
			}
		})
	}

}
