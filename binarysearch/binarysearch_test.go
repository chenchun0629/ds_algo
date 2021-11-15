package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a []int
	a = []int{1, 4, 7, 9, 12, 15, 16, 23, 35, 56, 67, 78, 90, 99}
	fmt.Println(BinarySearch(a, 9))
	fmt.Println(BinarySearch(a, 35))
}
