package sort

func MergeSort(data []int) int {
	return mergeSort(data, 0, len(data)-1)
}

func mergeSort(data []int, start, end int) (cnt int) {
	// 递归终止条件
	if start >= end {
		return cnt
	}

	var mid = (start + end) / 2
	cnt += mergeSort(data, start, mid)
	cnt += mergeSort(data, mid+1, end)

	cnt += merge(data, start, mid, end)

	return cnt
}

func merge(data []int, start, mid, end int) (cnt int) {
	var (
		tmp = make([]int, end-start+1)
		i   = start
		j   = mid + 1
		k   = 0
	)

	for i <= mid && j <= end {
		cnt++
		if data[i] <= data[j] {
			tmp[k] = data[i]
			i++
			k++
		} else {
			tmp[k] = data[j]
			j++
			k++
		}
	}

	var s = i
	var e = mid
	if j <= end {
		s = j
		e = end
	}

	for s <= e {
		tmp[k] = data[s]
		k++
		s++
	}

	for i := 0; i <= end-start; i++ {
		data[start+i] = tmp[i]
	}

	return cnt
}
