package singlelinkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList_NewSingleLinkedListFromSlice(t *testing.T) {
	var (
		s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		l = NewSingleLinkedListFromSlice(s)
	)

	l.Print()
}

func TestSingleLinkedList_Print(t *testing.T) {
	var (
		l  = NewSingleLinkedList(nil)
		_  = l.Append(0)
		n1  = l.Append(1)
		n2 = l.Append(2)
		_  = l.Append(3)
	)

	l.InsertBefore(n2, 1.5)
	fmt.Println(l.Find(1))
	fmt.Println(l.Remove(n2))
	l.MoveNodeToHead(n1)

	l.Print()

}
