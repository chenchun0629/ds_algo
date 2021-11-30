package consistent_hashing

import (
	"fmt"
	"testing"
)

func TestConsistent(t *testing.T) {
	var c = NewConsistent(20)
	c.Add("127.0.0.1:8080")
	c.Add("127.0.0.1:8081")
	c.Add("127.0.0.1:8082")
	c.Add("127.0.0.1:8083")

	var (
		table = []string{
			"lock:order:id:%d",
		}
	)

	for _, v := range table {

		var m = make(map[string]int)
		for i := 0; i < 1000000; i++ {
			var node, err = c.Get(fmt.Sprintf(v, i))
			if err == nil {
				m[node]++
			}
		}
		fmt.Println(m)
	}

}
