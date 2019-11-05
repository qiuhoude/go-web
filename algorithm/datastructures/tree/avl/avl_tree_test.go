package avl

import "testing"

var intCmpFunc CompareFunc = func(v, nodeV interface{}) int {
	vv := v.(int)
	nv := nodeV.(int)
	if vv > nv {
		return 1
	} else if vv < nv {
		return -1
	} else {
		return 0
	}
}

func TestAVLTree_Add(t *testing.T) {
	tree := NewAVLTree(intCmpFunc)
	for i := 0; i < 10; i++ {
		tree.Add(i)
	}
	t.Log(tree)
	t.Log(IsBST(tree))
	t.Log(IsBalanced(tree))
}
