package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	var ht = NewHashTable()
	ht.Put(StringValue("abc"))
	ht.Put(StringValue("abd"))
	ht.Put(StringValue("abe"))
	ht.Put(StringValue("bbc"))
	ht.Put(StringValue("bbd"))
	ht.Put(StringValue("bbe"))
	ht.Put(StringValue("cbc"))
	ht.Put(StringValue("cbd"))
	ht.Put(StringValue("cbe"))

	ht.Print()
}
