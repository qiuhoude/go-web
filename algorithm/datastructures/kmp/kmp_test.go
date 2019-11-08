package kmp

import (
	"testing"
)

func TestMarchSubstr(t *testing.T) {
	tests := []struct {
		mainStr  string
		subStr   string
		expected int
	}{
		{"BBC ABCDAB ABCDABCDABDEsddABCDABCDABDABCDABD", "ABCDABD", 3},
		{"BBC ABCDAB ABCDABCDABDE", "ABCDABD", 1},
		{"BBC ABCDAB ABCDABCDABDESS", "ABCDABD", 1},
		{"AAAAAAAAAAAAAAAAAAAAAAAA", "BB你好", 0},
	}
	for i, tt := range tests {
		out := MarchSubstr([]rune(tt.mainStr), []rune(tt.subStr))
		if out != tt.expected {
			t.Errorf("MarchSubstr() index=%d output %v, want %v", i, out, tt.expected)
		}
	}

}

func Test_prefixTable(t *testing.T) {
	s := []rune("ababacd")
	table1 := prefixTable(s)
	table2 := prefixTable2(s)
	t.Log(table1)
	t.Log(table2)
}
