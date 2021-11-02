package sample

import (
	"ds_algo/singlelinkedlist"
	"sync"
)

func NewLruCache(cap int) *LruCache {
	return &LruCache{
		cap:  cap,
		list: singlelinkedlist.NewSingleLinkedList(nil),
		rl:   sync.RWMutex{},
	}
}

type LruCache struct {
	cap int // 容量
	list *singlelinkedlist.SingleLinkedList // 单向列表

	rl sync.RWMutex // 读写锁
}

func (c *LruCache) Put(v interface{}) {
	c.rl.Lock()
	defer c.rl.Unlock()

	c.list.Prepend(v)

	for c.list.GetLength() > c.cap {
		c.list.Remove(c.list.Last())
	}
}

func (c *LruCache) Get(v interface{}) interface{} {
	c.rl.Lock()
	defer c.rl.Unlock()

	var n = c.list.Find(v)
	if n == nil {
		return nil
	}

	c.list.MoveNodeToHead(n)


	return n.GetValue()
}


func (c *LruCache) Print() {
	c.rl.RLock()
	defer c.rl.RUnlock()

	c.list.Print()
}


