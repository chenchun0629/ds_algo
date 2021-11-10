package sort

import (
	"sort"
)

func BubbleSort(data sort.Interface) int {
	return bubbleSort(data, data.Len())
}

func bubbleSort(data sort.Interface, l int) int {
	if l <= 1 {
		return 0
	}

	var cnt = 0
	for i := 0; i < l; i++ {
		var changed bool
		for j := i + 1; j < l; j++ {
			cnt++
			if data.Less(j, i) {
				changed = true
				data.Swap(j, i)
			}
		}

		if !changed {
			break
		}
	}

	return cnt
}
