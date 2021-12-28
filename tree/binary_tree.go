package tree

import (
	"ds_algo/stack"
	"fmt"
)

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(rootV interface{}) *BinaryTree {
	return &BinaryTree{NewNode(rootV)}
}

//四种主要的遍历思想为：
//
//前序遍历：根结点 ---> 左子树 ---> 右子树（深度优先遍历）
//
//中序遍历：左子树---> 根结点 ---> 右子树（深度优先遍历）
//
//后序遍历：左子树 ---> 右子树 ---> 根结点（深度优先遍历）
//
//层次遍历：只需按层次遍历即可（广度优先遍历）

// 中序遍历
func (this *BinaryTree) InOrderTraverse() {
	p := this.root

	s := stack.NewStackBaseList(99999)

	for !s.IsEmpty() || nil != p {
		if nil != p {
			s.Push(p)
			p = p.left
		} else {
			tmp := s.Pop().(*Node)
			fmt.Printf("%+v ", tmp.data)
			p = tmp.right
		}
	}
	fmt.Println()
}

// 前序遍历
func (this *BinaryTree) PreOrderTraverse() {
	p := this.root
	s := stack.NewStackBaseList(99999)

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.data)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*Node).right
		}
	}

	fmt.Println()
}

// 后序遍历
func (this *BinaryTree) PostOrderTraverse() {
	s1 := stack.NewStackBaseList(99999)
	s2 := stack.NewStackBaseList(99999)
	s1.Push(this.root)
	for !s1.IsEmpty() {
		p := s1.Pop().(*Node)
		s2.Push(p)
		if nil != p.left {
			s1.Push(p.left)
		}
		if nil != p.right {
			s1.Push(p.right)
		}
	}

	for !s2.IsEmpty() {
		fmt.Printf("%+v ", s2.Pop().(*Node).data)
	}

	fmt.Println()
}

// 广度优先
func (this *BinaryTree) BreadthFirstSearch() {
	var nodes = make([]*Node, 0)
	nodes = append(nodes, this.root)

	for len(nodes) > 0 {
		p := nodes[0]
		nodes = nodes[1:]
		fmt.Printf("%d ", p.data)

		if p.left != nil {
			nodes = append(nodes, p.left)
		}
		if p.right != nil {
			nodes = append(nodes, p.right)
		}
	}
	fmt.Println()
}

//use one stack, pre cursor to traverse from post order
func (this *BinaryTree) PostOrderTraverse2() {
	r := this.root
	s := stack.NewStackBaseList(99999)

	//point to last visit node
	var pre *Node

	s.Push(r)

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.left == nil && r.right == nil) ||
			(pre != nil && (pre == r.left || pre == r.right)) {

			fmt.Printf("%+v ", r.data)
			s.Pop()
			pre = r
		} else {
			if r.right != nil {
				s.Push(r.right)
			}

			if r.left != nil {
				s.Push(r.left)
			}

		}

	}
}
