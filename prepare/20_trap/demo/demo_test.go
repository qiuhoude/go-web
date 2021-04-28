package demo

import "testing"

func Test_slice1(t *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	t.Log(s, data)
	t.Log(&s[0], &data[0])

}
