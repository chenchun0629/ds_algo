package singlelinkedlist

import "reflect"

// ListNode 链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{value: v}
}

func (n *ListNode) GetNext() *ListNode {
	return n.next
}

func (n *ListNode) GetValue() interface{} {
	return n.value
}

func (n ListNode) CompareNode(c *ListNode) bool {
	if c == nil {
		return false
	}

	return n.CompareValue(c.value)
}

func (n ListNode) CompareValue(v interface{}) bool {
	var (
		ca, aok = n.value.(Comparable)
		cb, bok = v.(Comparable)
	)

	if aok && bok {
		return ca.IsEqual(cb)
	}

	return reflect.DeepEqual(n.value, v)
}