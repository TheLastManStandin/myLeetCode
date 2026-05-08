package main

import (
	"fmt"
	"math"
)

type Stack struct {
	topNode *Node
}

type Node struct {
	prev *Node
	val  int
	min  int
}

func MinStack() *Stack {
	headNode := &Node{
		prev: nil,
		val:  0,
		min:  int(math.Inf(1)),
	}
	stack := &Stack{
		topNode: headNode,
	}
	return stack
}

func (this *Stack) push(val int) {
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

func (this *Stack) pop() {
	this.topNode = this.topNode.prev
}

func (this *Stack) top() int {
	return this.topNode.val
}

func (this *Stack) getMin() int {
	return this.topNode.min
}

func main() {
	stack := MinStack()

	stack.push(1)
	stack.push(2)
	stack.push(0)
	fmt.Println(stack.getMin())
	stack.pop()
	fmt.Println(stack.top())
	fmt.Println(stack.getMin())
}
