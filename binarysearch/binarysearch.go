package binarysearch

func BinarySearch(data Searchable, t int) int {
	var l = data.Len()
	if l <= 0 {
		return -1
	}

	return binarySearch(data, 0, l-1, t)
}

func binarySearch(data Searchable, b, e, t int) int {
	if b > e {
		return -1
	}

	var mid = (b + e) / 2
	var c = data.Compare(mid, t)
	if c == 0 {
		return mid
	} else if c > 0 {
		return binarySearch(data, mid+1, e, t)
	}

	return binarySearch(data, b, mid-1, t)
}
