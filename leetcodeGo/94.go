package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rec(node *TreeNode) []int {
	if node.Left == nil && node.Right == nil {
		ret := make([]int, 0)
		ret = append(ret, node.Val)
		return ret
	}
	left := make([]int, 0)
	right := make([]int, 0)

	if node.Left != nil {
		left = append(left, rec(node.Left)...)
	}
	if node.Right != nil {
		right = append(right, rec(node.Right)...)
	}

	recAns := append(left, node.Val)
	recAns = append(recAns, right...)
	return recAns
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	ans := rec(root)

	return ans
}

func main() {
	two := &TreeNode{Val: 3, Left: nil, Right: nil}
	one := &TreeNode{
		Val:   2,
		Left:  two,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: one,
	}

	fmt.Println(inorderTraversal(root))
}
