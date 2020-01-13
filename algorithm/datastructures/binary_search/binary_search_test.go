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
	first2 := BSearchFirst2(arr, 8)
	t.Log("first2:", first2)
	last := BSearchLast(arr, 8)
	t.Log("last:", last)
	// 第一个大于等于
	fge := BSearchFirstGeVal(arr, 8)
	t.Log("fge:", fge)

	fg := BSearchFirstGVal(arr, 3)
	t.Log("fg:", fg)

	lle := BSearchLastLeVal(arr, 8)
	t.Log("lle:", lle)

}

func TestBSqrt(t *testing.T) {
	ret := BSqrt(3, 1e-6)
	t.Log(ret)
	ret2 := BSqrt(0.04, 1e-6)
	t.Log(ret2)
}
