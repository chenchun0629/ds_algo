package stack

import (
	"fmt"
	"testing"
)

func TestStackBaseList(t *testing.T) {
	//l := list.New()
	//_ = l.PushFront(4)
	//_ = l.PushFront(3)
	//_ = l.PushFront(2)
	//_ = l.PushFront(1)
	////l.InsertBefore(3, e4)
	////l.InsertAfter(2, e1)
	//
	//// Iterate through list and print its contents.
	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	var stack = NewStackBaseList(3)
	fmt.Println(stack.Push(1))
	fmt.Println(stack.Push(2))
	fmt.Println(stack.Push(3))
	fmt.Println(stack.Push(4))
	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Print()
}
