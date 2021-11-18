package skiplist

type Nodable interface {
	Weight() int
	Compare(v Nodable) int
}

type IntNode int

func (i IntNode) Compare(v Nodable) int {
	if d, ok := v.(IntNode); ok {
		if i > d {
			return 1
		} else if i == d {
			return 0
		}

		return -1
	}

	panic("Nodable v's type is not IntNode")
}

func (i IntNode) Weight() int {
	return int(i)
}
