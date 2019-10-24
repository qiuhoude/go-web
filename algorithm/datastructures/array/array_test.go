package array

import "testing"

func loadArry() (*Array, error) {
	capacity := 10
	arr := New(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			return nil, err
		}
	}
	return arr, nil
}

func TestArray_Insert(t *testing.T) {
	arr, err := loadArry()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)

	arr.Insert(uint(6), 999)
	t.Log(arr)

	arr.InsertToTail(666)
	t.Log(arr)
}

func TestArray_Delete(t *testing.T) {
	arr, err := loadArry()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)
	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		t.Log(arr)
	}
}

func TestArray_Find(t *testing.T) {
	arr, err := loadArry()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)

	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}
