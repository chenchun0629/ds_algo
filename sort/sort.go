package sort

import (
	"sort"
)

type Interface interface {
	sort.Interface

	Divide() (Interface, Interface)
	Merge(q Interface) Interface
}

type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x IntSlice) Divide() (Interface, Interface) {
	var hl = len(x) / 2
	var q = (x)[hl:]
	x = (x)[:hl:hl]
	return x, q
}

func (x IntSlice) Merge(q Interface) Interface {
	return append(x, q.(IntSlice)...)
}
