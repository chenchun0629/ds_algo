package binarysearch

func BinarySearch(data []int, t int) int {
	var l = len(data)
	if l <= 0 {
		return -1
	}

	return binarySearch(data, 0, l-1, t)
}

func binarySearch(data []int, b, e, t int) int {
	if b > e {
		return -1
	}

	var mid = (b + e) / 2
	if data[mid] == t {
		return mid
	} else if t > data[mid] {
		return binarySearch(data, mid+1, e, t)
	}

	return binarySearch(data, b, mid-1, t)
}
