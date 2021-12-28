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
	var (
		inOrder func(p *Node)
	)

	inOrder = func(p *Node) {
		if p == nil {
			return
		}

		inOrder(p.left)
		fmt.Print(p.data, " ")
		inOrder(p.right)
	}

	inOrder(this.root)

	fmt.Println()
}

// 前序遍历
func (this *BinaryTree) PreOrderTraverse() {
	var (
		preOrder func(p *Node)
	)

	preOrder = func(p *Node) {
		if p == nil {
			return
		}

		fmt.Print(p.data, " ")
		preOrder(p.left)
		preOrder(p.right)
	}

	preOrder(this.root)
	fmt.Println()
}

// 后序遍历
func (this *BinaryTree) PostOrderTraverse() {
	var (
		postOrder func(p *Node)
	)

	postOrder = func(p *Node) {
		if p == nil {
			return
		}

		postOrder(p.left)
		postOrder(p.right)
		fmt.Print(p.data, " ")
	}

	postOrder(this.root)
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
