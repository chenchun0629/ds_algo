package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	var sl = NewSkipList()

	for i := 0; i < 10000; i++ {
		sl.Set(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 2 {
		sl.Remove(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 3 {
		sl.Remove(i)
	}

	sl.Print()
	for i := 0; i < 10000; i += 4 {
		sl.Remove(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 5 {
		sl.Remove(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 7 {
		sl.Remove(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 11 {
		sl.Remove(i)
	}

	sl.Print()

	for i := 0; i < 10000; i += 13 {
		sl.Remove(i)
	}

	sl.Print()
}
