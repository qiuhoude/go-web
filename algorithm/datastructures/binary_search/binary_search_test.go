package binary_search

import "testing"

func TestBSearch(t *testing.T) {
	arr := []int{1, 3, 4, 5, 6, 6, 8, 8, 8, 11, 18}
	i := BSearch(arr, 8)
	t.Log(i)
	i2 := BSearchRecusion(arr, 8)
	t.Log(i2)
	first := BSearchFirst(arr, 8)
	t.Log("first:", first)
	last := BSearchLast(arr, 8)
	t.Log("last:", last)
	// 第一个大于等于
	fge := BSearchFirstGeVal(arr, 18)
	t.Log("fge:", fge)

	lle := BSearchLastLeVal(arr, 18)
	t.Log("lle:", lle)

}

func TestBSqrt(t *testing.T) {
	ret := BSqrt(3, 1e-6)
	t.Log(ret)
	ret2 := BSqrt(0.04, 1e-6)
	t.Log(ret2)
}
