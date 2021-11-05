package queue

import (
	"fmt"
	"testing"
)

func TestQueueBaseSlice(t *testing.T) {
	var stack = NewQueueBaseSlice(3)
	fmt.Println(stack.Enqueue(1))
	fmt.Println(stack.Enqueue(2))
	fmt.Println(stack.Enqueue(3))
	fmt.Println(stack.Enqueue(4))
	stack.Print()
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	stack.Print()
}
