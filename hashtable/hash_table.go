package hashtable

import "fmt"

type HashTable struct {
	m map[uint8]*Node
}

func NewHashTable() *HashTable {
	return &HashTable{
		m: make(map[uint8]*Node),
	}
}

func (t *HashTable) Put(v Valuable) bool {
	var (
		key = v.GetKey()
		hk  = t.getHashCode(key)
	)

	if _, has := t.m[hk]; !has {
		t.m[hk] = NewNode(fmt.Sprintf("%d", hk))
	}

	return t.m[hk].Push(v)
}

func (t *HashTable) Get(key string) Valuable {
	var hk = t.getHashCode(key)

	if _, has := t.m[hk]; !has {
		return nil
	}

	return t.m[hk].Find(key)
}

func (t *HashTable) Remove(key string) bool {
	var hk = t.getHashCode(key)

	if _, has := t.m[hk]; !has {
		return true
	}

	return t.m[hk].Remove(key)
}

func (t *HashTable) Print() {
	for _, v := range t.m {
		v.Print()
	}
}

// 一个简单容易碰撞的hash算法
func (t HashTable) getHashCode(s string) uint8 {
	if len(s) == 0 {
		return 0
	}

	return s[0]
}
