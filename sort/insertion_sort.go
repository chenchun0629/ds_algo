package sort

import (
	"sort"
)

func InsertionSort(data sort.Interface) int {
	return insertionSort(data, data.Len())
}

func insertionSort(data sort.Interface, l int) int {
	if l <= 1 {
		return 0
	}

	var cnt = 0
	for i := 1; i < l; i++ {
		k := i
		for j := k - 1; j >= 0; j-- {
			cnt++
			if data.Less(k, j) {
				data.Swap(k, j)
				k--
			} else {
				break
			}
		}
	}

	return cnt
}
