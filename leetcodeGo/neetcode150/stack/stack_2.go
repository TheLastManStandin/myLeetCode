package stack

import "math"

type MinStack struct {
	topNode *Node
}

type Node struct {
	prev *Node
	val  int
	min  int
}

func Constructor() MinStack {
	headNode := &Node{
		prev: nil,
		val:  0,
		min:  int(math.Pow(2, 31)),
	}
	stack := MinStack{
		topNode: headNode,
	}
	return stack
}

func (this *MinStack) Push(val int) {
	newTop := Node{
		prev: this.topNode,
		val:  val,
	}

	if val < this.topNode.min {
		newTop.min = val
	} else {
		newTop.min = this.topNode.min
	}

	this.topNode = &newTop
}

func (this *MinStack) Pop() {
	this.topNode = this.topNode.prev
}

func (this *MinStack) Top() int {
	return this.topNode.val
}

func (this *MinStack) GetMin() int {
	return this.topNode.min
}
