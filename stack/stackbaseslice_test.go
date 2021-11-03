package stack

import (
	"fmt"
	"testing"
)

func TestStackBaseSlice_removeIndex(t *testing.T) {
	var stack = NewStackBaseSlice(3)
	fmt.Println(stack.Push(1))
	fmt.Println(stack.Push(2))
	fmt.Println(stack.Push(3))
	fmt.Println(stack.Push(4))
	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Print()
}
