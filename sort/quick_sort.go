package sort

func QuickSort(data Interface) int {
	return quickSort(data, 0, data.Len()-1)
}

func quickSort(data Interface, start, end int) int {
	// 终止递归条件
	if start >= end {
		return 0
	}

	// 进行分区排序，并确定中点
	var mid, cnt = partition(data, start, end)

	// 递归
	cnt += quickSort(data, start, mid-1)
	cnt += quickSort(data, mid+1, end)

	return cnt
}

func partition(data Interface, start, end int) (mid int, cnt int) {
	// 设中间点mid为起点， 并且 拿最后一个值data[end]为标杆 进行比较
	// 从start开始遍历  与标杆进行data[end]比较
	// 如果 当前值 data[start] < 标杆 data[end]，
	// 那么 操作 data[mid] 中点值 与 data[start] 开始点值 交换，并且 中点index 向后移动一位
	// 最后 将 mid 和 end 标杆 的 值进行交换

	// mid++代表了什么？ 当 当前值 data[start] 小于标杆值data[end]  时 mid++，
	// 说明了 小于标杆值的数量总共有mid个，那么最后mid和end交换，其实就是把标杆值放到正确的位置。

	// data.Swap(mid, start) 代表了什么？ 当 当前值 data[start] 小于标杆值data[end] 时 data.Swap(mid, start)，
	// 说明了 data[start]的值小于 data[mid]的值，两个值交换，就是把大数后移

	mid = start
	for ; start <= end-1; start++ {
		cnt++
		if data.Less(start, end) {
			if mid != start {
				data.Swap(mid, start)
			}

			mid++
		}
	}

	data.Swap(mid, end)

	return mid, cnt
}
