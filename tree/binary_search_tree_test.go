package tree

import "testing"

func TestBST_Find(t *testing.T) {
	bst := NewBST(1, IntCompareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)
	bst.PreOrderTraverse()
	t.Log(bst.Find(2))
	t.Log(bst.Min())
	t.Log(bst.Max())
	t.Log(bst.Delete(7))
	bst.InOrderTraverse()
}
