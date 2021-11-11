package sort

import "sort"

type Interface interface {
	sort.Interface

	Get(i int) interface{}
	Set(i int, v interface{})
}

type IntSlice []int

func (x IntSlice) Set(i int, v interface{}) {
	x[i] = v.(int)
}

func (x IntSlice) Get(i int) interface{} {
	return x[i]
}

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
