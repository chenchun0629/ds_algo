package binarysearch

type Searchable interface {
	Compare(index int, target interface{}) int
	Len() int
	Get(index int) interface{}
}

type IntSlice []int

func (i IntSlice) Compare(index int, target interface{}) int {
	if data, ok := target.(int); ok {
		if data == i[index] {
			return 0
		} else if data > i[index] {
			return 1
		}
		return -1
	}

	panic("invalid target value type")
}

func (i IntSlice) Len() int {
	return len(i)
}

func (i IntSlice) Get(index int) interface{} {
	return i[index]
}
