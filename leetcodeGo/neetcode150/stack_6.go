package main

type doubleStack [][2]int

func (s *doubleStack) Push(v [2]int) {
	*s = append(*s, v)
}

func (s *doubleStack) Pop() [2]int {
	lens := len(*s)
	res := (*s)[lens-1]
	*s = (*s)[:lens-1]
	return res
}

func largestRectangleArea(heights []int) int {
	stk := doubleStack{}
	ans := 0

	for i, h := range heights {

	}
}
