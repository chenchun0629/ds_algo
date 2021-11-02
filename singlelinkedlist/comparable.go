package singlelinkedlist

type Comparable interface {
	IsEqual(v Comparable) bool
	GetCompareData() interface{}
}
