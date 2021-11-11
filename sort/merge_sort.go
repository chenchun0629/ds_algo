package sort

func MergeSort(data Interface) int {
	return mergeSort(data, 0, data.Len()-1)
}

func mergeSort(data Interface, start, end int) (cnt int) {
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

func merge(data Interface, start, mid, end int) (cnt int) {
	var (
		tmp = make([]interface{}, end-start+1)
		i   = start
		j   = mid + 1
		k   = 0
	)

	for i <= mid && j <= end {
		cnt++
		if data.Less(i, j) {
			tmp[k] = data.Get(i)
			i++
			k++
		} else {
			tmp[k] = data.Get(j)
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
		tmp[k] = data.Get(s)
		k++
		s++
	}

	for i := 0; i <= end-start; i++ {
		data.Set(start+i, tmp[i])
	}

	return cnt
}
