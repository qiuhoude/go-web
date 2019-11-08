package bm

import (
	"fmt"
	"testing"
)

func Test_generateGS(t *testing.T) {
	suffix, prefix := generateGS([]rune("cabcabcab"))
	t.Log(suffix)
	t.Log(prefix)
}

func Test_BmSearch(t *testing.T) {
	main := "abcacabcb洗cabcab哈哈c"
	pattern := "cabcab"
	fmt.Println(BmSearch([]rune(main), []rune(pattern)))
}
