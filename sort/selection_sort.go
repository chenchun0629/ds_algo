package sort

import "sort"

func SelectionSort(data sort.Interface) int {
	return selectionSort(data, data.Len())
}

func selectionSort(data sort.Interface, l int) int {
	if l <= 1 {
		return 0
	}

	var cnt = 0

	for i := 1; i < l; i++ {
		var min = i - 1
		for j := i; j < l; j++ {
			cnt++
			if data.Less(j, min) {
				min = j
			}
		}

		if i-1 != min {
			data.Swap(i-1, min)
		}
	}

	return cnt
}
