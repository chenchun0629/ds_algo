package stack

import "fmt"

func NewStackBaseSlice(cap int) *StackBaseSlice {
	return &StackBaseSlice{items: make([]interface{}, 0, cap), cap: cap}
}

type StackBaseSlice struct {
	items []interface{}
	len   int
	cap   int
}

func (s *StackBaseSlice) Push(v interface{}) bool {
	if s.len >= s.cap {
		return false
	}

	s.insertValue(0, v)
	return true
}

func (s *StackBaseSlice) Pop() interface{} {
	if s.len == 0 {
		return nil
	}

	return s.removeIndex(0)
}

func (s *StackBaseSlice) removeIndex(i int) interface{} {
	s.len--
	var v = s.items[i]
	s.items = append(s.items[:i], s.items[i+1:]...)
	return v
}

func (s *StackBaseSlice) insertValue(i int, v interface{}) {
	s.len++
	s.items = append(s.items, nil)
	copy(s.items[i+1:], s.items[i:])
	s.items[i] = v
}

func (s *StackBaseSlice) Print() {
	fmt.Printf("{items: %#v, len: %d, cap: %d} \n", s.items, s.len, s.cap)
}
