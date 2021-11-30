package consistent_hashing

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type uints []uint32

func (x uints) Len() int           { return len(x) }
func (x uints) Less(i, j int) bool { return x[i] < x[j] }
func (x uints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

var ErrEmptyNodes = errors.New("empty nodes")

type Consistent struct {
	replicas int

	circle       map[uint32]string
	nodes        map[string]bool
	sortedHashes uints

	sync.RWMutex
}

func NewConsistent(replicas int) *Consistent {
	return &Consistent{
		replicas: replicas,
		circle:   make(map[uint32]string),
		nodes:    make(map[string]bool),
	}
}

func (c *Consistent) Get(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.circle) == 0 {
		return "", ErrEmptyNodes
	}

	var (
		h  = c.hashKey(key)
		ci = c.search(h)
	)

	return c.circle[c.sortedHashes[ci]], nil
}

func (c *Consistent) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func (c *Consistent) Add(nodeName string) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.replicas; i++ {
		c.circle[c.hashKey(c.genNodeKey(nodeName, i))] = nodeName
	}

	c.updateSortedHashes()
	c.nodes[nodeName] = true
}

func (c *Consistent) Remove(nodeName string) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.replicas; i++ {
		delete(c.circle, c.hashKey(c.genNodeKey(nodeName, i)))
	}
	delete(c.nodes, nodeName)
	c.updateSortedHashes()
}

func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *Consistent) genNodeKey(node string, index int) string {
	return fmt.Sprintf("%s#%d", node, index)
}

func (c *Consistent) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(c.sortedHashes)/(c.replicas*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
}
