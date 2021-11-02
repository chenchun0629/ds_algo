package singlelinkedlist

import (
	"fmt"
	"reflect"
)

func NewSingleLinkedListFromSlice(v interface{}) *SingleLinkedList {
	var (
		rv = reflect.ValueOf(v)
		l  = NewSingleLinkedList(nil)
	)

	switch rv.Kind() {
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			l.Append(rv.Index(i).Interface())
		}
	default:
		panic("v is not a slice")
	}

	return l
}

func NewSingleLinkedList(head *ListNode) *SingleLinkedList {
	var l = &SingleLinkedList{head: head}
	if head != nil {
		l.nodeLengthIncr()
	}
	return l
}

// SingleLinkedList 单向链表
type SingleLinkedList struct {
	head   *ListNode
	length int
}

func (l SingleLinkedList) Print() {
	var c = l.head
	for c != nil {
		fmt.Printf("%#v \n", c.value)
		c = c.next
	}
}

//Prepend 添加至头
func (l *SingleLinkedList) Prepend(v interface{}) *ListNode {
	var n = NewListNode(v)
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head = n
	}

	l.nodeLengthIncr()
	return n
}

func (l *SingleLinkedList) MoveNodeToHead(n *ListNode) bool {
	if n == nil {
		return false
	}

	var (
		prev *ListNode
		c    = l.head
	)
	for c != nil {
		if c == n {
			if prev != nil {
				prev.next = n.next
				n.next = l.head
				l.head = n
			}
			return true
		}

		prev = c
		c = c.next
	}

	return false
}

//Append 添加至最后
func (l *SingleLinkedList) Append(v interface{}) *ListNode {
	var n = NewListNode(v)
	if l.head == nil {
		l.head = n
	} else {
		var (
			c    = l.head
			last *ListNode
		)

		for c != nil {
			last = c
			c = c.next
		}

		last.next = n
	}

	l.nodeLengthIncr()
	return n

}

//Insert 添加至某个节点之后
func (l *SingleLinkedList) Insert(p *ListNode, v interface{}) *ListNode {
	if p == nil {
		return nil
	}

	var c = NewListNode(v)
	c.next = p.next
	p.next = c

	l.nodeLengthIncr()

	return c
}

func (l *SingleLinkedList) InsertBefore(p *ListNode, v interface{}) *ListNode {
	if p == nil {
		return nil
	}

	var (
		prev *ListNode
		c    = l.head
		n    = NewListNode(v)
	)
	for c != nil {
		if c == p {
			if prev == nil {
				l.head = n
				n.next = p
			} else {
				prev.next = n
				n.next = p
			}
			l.nodeLengthIncr()
			return n
		}

		prev = c
		c = c.next
	}

	return nil
}

func (l *SingleLinkedList) Remove(n *ListNode) bool {
	if n == nil {
		return true
	}

	var (
		prev *ListNode
		c    = l.head
	)
	for c != nil {
		if c == n {
			if prev == nil {
				l.head = c.next
			} else {
				prev.next = c.next
			}
			l.nodeLengthDecr()
			return true
		}

		prev = c
		c = c.next
	}

	return false
}

func (l SingleLinkedList) Last() *ListNode {
	var c = l.head
	if c == nil {
		return nil
	}

	for c != nil {
		if c.next == nil {
			return c
		}
		c = c.next
	}

	return nil
}

func (l *SingleLinkedList) Find(v interface{}) *ListNode {
	var c = l.head
	for c != nil {
		if c.CompareValue(v) {
			return c
		}
		c = c.next
	}

	return nil
}

func (l *SingleLinkedList) nodeLengthIncr() {
	l.length++
}
func (l *SingleLinkedList) nodeLengthDecr() {
	l.length--
}

func (l SingleLinkedList) GetLength() int {
	return l.length
}
