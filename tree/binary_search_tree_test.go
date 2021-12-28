package tree

import (
	"fmt"
	"testing"
)

func TestBST_Find(t *testing.T) {
	bst := NewBST(1, IntCompareFunc)

	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(7)
	fmt.Printf("bfs: ")
	bst.BreadthFirstSearch()
	fmt.Printf("preot: ")
	bst.PreOrderTraverse()
	fmt.Printf("postot: ")
	bst.PostOrderTraverse()
	fmt.Printf("inot: ")
	bst.InOrderTraverse()
	t.Log(bst.Find(2))
	t.Log(bst.Min())
	t.Log(bst.Max())
	t.Log(bst.Delete(7))
}
