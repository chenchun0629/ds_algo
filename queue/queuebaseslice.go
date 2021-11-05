package queue

import "fmt"

func NewQueueBaseSlice(cap int) *QueueBaseSlice {
	return &QueueBaseSlice{
		items: make([]interface{}, 0),
		cap:   cap,
		len:   0,
	}
}

type QueueBaseSlice struct {
	items []interface{}
	cap   int
	len   int
}

func (s *QueueBaseSlice) Enqueue(v interface{}) bool {
	if s.len >= s.cap {
		return false
	}

	s.items = append(s.items, v)
	s.len++
	return true
}

func (s *QueueBaseSlice) Dequeue() interface{} {
	if s.len <= 0 {
		return nil
	}

	var v = s.items[0]
	s.items = s.items[1:]
	s.len--
	return v
}

func (s QueueBaseSlice) Print() {
	fmt.Printf("{items: %#v, len: %d, cap: %d} \n", s.items, s.len, s.cap)
}
