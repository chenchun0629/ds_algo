package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var nums = []int{345, 425, 2, 457, 346, 735, 68, 356, 256, 2, 756, 83, 7946, 3, 4}
	//var nums = []int{2}
	fmt.Println(BubbleSort(sort.IntSlice(nums)), nums)
}
