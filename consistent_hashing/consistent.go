package consistent_hashing

import (
	"errors"
	"fmt"
	"hash/crc32"
	"math"
	"sort"
	"sync"
	"sync/atomic"
)

type uints []uint32

func (x uints) Len() int           { return len(x) }
func (x uints) Less(i, j int) bool { return x[i] < x[j] }
func (x uints) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

var ErrEmptyNodes = errors.New("empty nodes")

type options struct {
	replicas int // 副本/虚拟节点数量

	loadFactor float64 // 负载因子
}

func (o *options) apply(fn ...SetOptionsFunc) {
	for _, v := range fn {
		v(o)
	}
}

func newOptions() *options {
	return &options{
		replicas:   20,
		loadFactor: 1.25,
	}
}

type SetOptionsFunc func(*options)

func SetOptionsReplicas(replicas int) SetOptionsFunc {
	return func(o *options) {
		o.replicas = replicas
	}
}
func SetOptionsLoadFactor(loadFactor float64) SetOptionsFunc {
	return func(o *options) {
		o.loadFactor = loadFactor
	}
}

type Node struct {
	Name string
	Load int64
}

type Consistent struct {
	options *options

	circle       map[uint32]string
	nodes        map[string]*Node
	sortedHashes uints

	totalLoad int64

	sync.RWMutex
}

func NewConsistent(fn ...SetOptionsFunc) *Consistent {
	var o = newOptions()
	o.apply(fn...)

	return &Consistent{
		options:   o,
		circle:    make(map[uint32]string),
		nodes:     make(map[string]*Node),
		totalLoad: 0,
	}
}

func (c *Consistent) Incr(node string) {
	c.Lock()
	defer c.Unlock()

	atomic.AddInt64(&c.nodes[node].Load, 1)
	atomic.AddInt64(&c.totalLoad, 1)
}

func (c *Consistent) Done(node string) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.nodes[node]; !ok {
		return
	}
	atomic.AddInt64(&c.nodes[node].Load, -1)
	atomic.AddInt64(&c.totalLoad, -1)
}

func (c *Consistent) GetLeast(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.nodes) == 0 {
		return "", ErrEmptyNodes
	}

	var (
		h   = c.hashKey(key)
		idx = c.search(h)
		i   = idx
	)

	for {
		var node = c.circle[c.sortedHashes[i]]
		if c.loadNotOver(node) {
			return node, nil
		}
		i++
		if i >= len(c.circle) {
			i = 0
		}

		if i == idx {
			break
		}
	}

	return c.circle[c.sortedHashes[idx]], nil
}

func (c *Consistent) loadNotOver(node string) bool {
	if c.totalLoad < 0 {
		c.totalLoad = 0
	}

	var avgLoad = float64((c.totalLoad + 1) / int64(len(c.nodes)))
	avgLoad = math.Ceil(avgLoad * c.options.loadFactor)

	var n, has = c.nodes[node]
	if !has {
		panic(fmt.Sprintf("given node(%s) not in nodes", n.Name))
	}

	if float64(n.Load) <= avgLoad {
		return true
	}

	return false
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

func (c *Consistent) Add(key string) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.options.replicas; i++ {
		c.circle[c.hashKey(c.genNodeKey(key, i))] = key
	}

	c.updateSortedHashes()
	c.nodes[key] = &Node{
		Name: key,
		Load: 0,
	}
}

func (c *Consistent) Remove(key string) {
	c.Lock()
	defer c.Unlock()

	for i := 0; i < c.options.replicas; i++ {
		delete(c.circle, c.hashKey(c.genNodeKey(key, i)))
	}
	delete(c.nodes, key)
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
	if cap(c.sortedHashes)/(c.options.replicas*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
}
