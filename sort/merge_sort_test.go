package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var nums = []int{345, 425, 2, 457, 346, 735, 68, 356, 256, 2, 756, 83, 7946, 3, 4}
	//var nums = []int{1, 4, 2, 5, 3}
	//var nums = []int{2}
	fmt.Println(MergeSort(IntSlice(nums)), nums)
}
