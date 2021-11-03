package stack

import (
	"container/list"
	"fmt"
)

func NewStackBaseList(cap int) *StackBaseList {
	return &StackBaseList{cap: cap, list: list.New()}
}

type StackBaseList struct {
	list *list.List
	cap  int
}

func (l *StackBaseList) Push(v interface{}) bool {
	if l.list.Len() >= l.cap {
		return false
	}

	l.list.PushFront(v)
	return true
}

func (l *StackBaseList) Pop() interface{} {
	if l.list.Len() <= 0 {
		return nil
	}

	var e = l.list.Front()
	l.list.Remove(e)
	return e.Value
}

func (l *StackBaseList) Print() {
	var i = make([]interface{}, 0)
	for e := l.list.Front(); e != nil; e = e.Next() {
		i = append(i, e.Value)
	}
	fmt.Printf("{items: %#v, len: %d, cap: %d} \n", i, l.list.Len(), l.cap)
}
