package queue

import (
	"fmt"
	"testing"
)

func TestQueueBaseList(t *testing.T) {
	var stack = NewQueueBaseList(3)
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
