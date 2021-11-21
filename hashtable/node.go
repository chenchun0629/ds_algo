package hashtable

import "fmt"

type Node struct {
	hashKey string
	store   []Valuable
}

func (n *Node) Push(v Valuable) bool {
	for _, nv := range n.store {
		if nv.GetKey() == v.GetKey() {
			return false
		}
	}

	n.store = append(n.store, v)
	return true
}

func (n *Node) Find(key string) Valuable {
	for _, nv := range n.store {
		if nv.GetKey() == key {
			return nv
		}
	}

	return nil
}

func (n *Node) Remove(key string) bool {
	for nk, nv := range n.store {
		if nv.GetKey() == key {
			n.store = append(n.store[:nk], n.store[nk+1:]...)
			return true
		}
	}

	return true
}

func (n *Node) Print() {
	fmt.Printf("hash key: %s, data: %#v \n", n.hashKey, n.store)
}

func NewNode(hashKey string) *Node {
	return &Node{
		hashKey: hashKey,
		store:   make([]Valuable, 0),
	}
}
