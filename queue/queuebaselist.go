package queue

import (
	"container/list"
	"fmt"
)

func NewQueueBaseList(cap int) *QueueBaseList {
	return &QueueBaseList{
		list: list.New(),
		cap:  cap,
	}
}

type QueueBaseList struct {
	list *list.List
	cap  int
}

func (s *QueueBaseList) Enqueue(v interface{}) bool {
	if s.list.Len() >= s.cap {
		return false
	}

	s.list.PushBack(v)
	return true
}

func (s *QueueBaseList) Dequeue() interface{} {
	if s.list.Len() <= 0 {
		return nil
	}

	var v = s.list.Front()
	s.list.Remove(v)
	return v.Value
}

func (s QueueBaseList) Print() {
	var i = make([]interface{}, 0)
	for e := s.list.Front(); e != nil; e = e.Next() {
		i = append(i, e.Value)
	}
	fmt.Printf("{items: %#v, len: %d, cap: %d} \n", i, s.list.Len(), s.cap)
}
