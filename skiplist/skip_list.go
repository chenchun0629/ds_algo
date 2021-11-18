package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Value Nodable
	Next  []*Node
}

func NewNode(value Nodable, maxLevel int) *Node {
	return &Node{
		Value: value,
		Next:  make([]*Node, maxLevel),
	}
}

func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{
		head:            NewNode(nil, maxLevel),
		maxLevel:        maxLevel,
		currentMaxLevel: 1,
		rand:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type SkipList struct {
	head *Node

	maxLevel        int
	currentMaxLevel int
	rand            *rand.Rand
}

func (l *SkipList) Get(value Nodable) *Node {
	var p = l.head
	// 从上层开始遍历
	for i := l.currentMaxLevel; i >= 0; i-- {
		for ; p.Next[i] != nil && p.Next[i].Value.Compare(value) < 0; p = p.Next[i] {
		}
	}

	if p.Next[0] != nil && p.Next[0].Value == value {
		return p.Next[0]
	}
	return nil
}

func (l *SkipList) Remove(value Nodable) {
	var (
		update = make([]*Node, l.currentMaxLevel)
		p      = l.head
	)

	// 从上层开始遍历
	for i := l.currentMaxLevel - 1; i >= 0; i-- {
		for ; p.Next[i] != nil && p.Next[i].Value.Compare(value) < 0; p = p.Next[i] {
		}
		update[i] = p
	}

	// 移除关联关系
	for i := 0; i < l.currentMaxLevel; i++ {
		if update[i].Next[i] != nil && update[i].Next[i].Value == value {
			update[i].Next[i] = update[i].Next[i].Next[i]
		}
	}

	// 判断 currentMaxLevel 是否需要调整
	// 第一行不移除
	for i := l.currentMaxLevel - 1; i > 0; i-- {
		if l.head.Next[i] == nil {
			l.currentMaxLevel--
		}
	}

}

func (l *SkipList) Set(value Nodable) {
	var (
		lvl    = l.randomMaxLevel()  // 当前内容索引等级
		node   = NewNode(value, lvl) // 当前内容节点
		update = make([]*Node, lvl)  // 需要更新节点列表
	)

	for i := 0; i < lvl; i++ {
		update[i] = l.head // 默认需要更新的都是首节点
	}

	// 从最上层索引开始计算需要更新的节点
	// 这里主要为了找到第i层需要更新的节点，并且记录到update中去
	var p = l.head
	for i := lvl - 1; i >= 0; i-- {
		// 第i层索引不为空，并且第i层索引的值小于插入的值
		for p.Next[i] != nil && p.Next[i].Value.Compare(value) < 0 {
			p = p.Next[i] // 节点p就滑动到下一个
		}
		update[i] = p
	}

	// 正式关联节点
	for i := 0; i < lvl; i++ {
		// 将 node.next 指向待更新节点的next
		node.Next[i] = update[i].Next[i]
		// 将待更新节点的next，指向node
		update[i].Next[i] = node
	}

	if l.currentMaxLevel < lvl {
		l.currentMaxLevel = lvl
	}
}

// 理论来讲，一级索引中元素个数应该占原始数据的 50%，二级索引中元素个数占 25%，三级索引12.5% ，一直到最顶层。
// 因为这里每一层的晋升概率是 50%。对于每一个新插入的节点，都需要调用 randomLevel 生成一个合理的层数。
// 该 randomLevel 方法会随机生成 1~MAX_LEVEL 之间的数，且 ：
//        50%的概率返回 1
//        25%的概率返回 2
//      12.5%的概率返回 3 ...
func (l SkipList) randomMaxLevel() int {
	const prob = 1 << 30 // Half of 2^31.
	var (
		estimated = l.maxLevel
		rd        = l.rand
		lvl       = 1
	)

	for ; lvl < estimated; lvl++ {
		if rd.Int31() < prob {
			break
		}
	}

	return lvl
}

func (l SkipList) Print() {
	for i := 0; i < l.currentMaxLevel; i++ {
		fmt.Printf("第%d层： ", i)
		for p := l.head; p != nil; p = p.Next[i] {
			fmt.Printf("%#v \t", p.Value)
		}
		fmt.Println()
	}
}
