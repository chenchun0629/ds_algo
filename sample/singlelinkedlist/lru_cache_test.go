package singlelinkedlist

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	var c = NewLruCache(3)
	c.Put(0)
	c.Put(1)
	c.Put(2)
	c.Print()
	fmt.Println(c.Get(0))
	c.Put(5)
	c.Print()
}
