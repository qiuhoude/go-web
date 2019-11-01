package sikplist

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("leo", 95)
	sl.Insert("jack", 88)
	sl.Insert("lily", 100)
	sl.Insert("jack", 88)
	sl.Insert("jack1", 88)
	sl.Insert("jack", 83)
	sl.Insert("jack300", 300)
	t.Log(sl)
	t.Log("-----------------------------")

	t.Log(sl.Find("jack", 88))
	t.Log("-----------------------------")

	sl.Delete("leo", 95)

	t.Log(sl)
	t.Log("-----------------------------")
}
