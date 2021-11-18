package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	var sl = NewSkipList(16)

	for i := 0; i < 100; i++ {
		sl.Set(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 2 {
		sl.Remove(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 3 {
		sl.Remove(IntNode(i))
	}

	sl.Print()
	for i := 0; i < 100; i += 4 {
		sl.Remove(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 5 {
		sl.Remove(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 7 {
		sl.Remove(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 11 {
		sl.Remove(IntNode(i))
	}

	sl.Print()

	for i := 0; i < 100; i += 13 {
		sl.Remove(IntNode(i))
	}

	sl.Print()
}
