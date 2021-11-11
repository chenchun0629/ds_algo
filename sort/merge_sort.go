package sort

func MergeSort(data Interface) int {
	return mergeSort(data, 0, data.Len()-1)
}

func mergeSort(data Interface, start, end int) (cnt int) {
	// 递归终止条件
	if start >= end {
		return cnt
	}

	// 分治递归
	var mid = (start + end) / 2
	cnt += mergeSort(data, start, mid)
	cnt += mergeSort(data, mid+1, end)

	// 结果合并
	cnt += merge(data, start, mid, end)

	return cnt
}

func merge(data Interface, start, mid, end int) (cnt int) {

	// 开始比较 start->mid, mid->end 的数据
	// 有序的塞进tmp
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

	// 将剩余未塞入tmp的数据继续塞入
	// 先确定 start->mid 未执行完 还是 mid->end未执行完
	var s = i
	var e = mid
	if j <= end {
		s = j
		e = end
	}

	// 将剩余的数据塞入tmp
	for s <= e {
		tmp[k] = data.Get(s)
		k++
		s++
	}

	// 把tmp中的数据有序的合并进data
	for i := 0; i <= end-start; i++ {
		data.Set(start+i, tmp[i])
	}

	return cnt
}
