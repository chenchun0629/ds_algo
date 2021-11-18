package skiplist

type Nodable interface {
	Weight() int
	Compare(v Nodable) int
}

type IntNode int

func (i IntNode) Compare(v Nodable) int {
	if i.Weight() > v.Weight() {
		return 1
	} else if i.Weight() == v.Weight() {
		return 0
	}

	return -1

}

func (i IntNode) Weight() int {
	return int(i)
}
