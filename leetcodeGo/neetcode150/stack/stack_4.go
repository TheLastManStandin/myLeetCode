package stack

func dailyTemperatures(temperatures []int) []int {
	stack := Constructor()
	res := make([]int, len(temperatures))

	for i, t := range temperatures {
		if stack.topNode.prev == nil {
			stack.Push(i)
		} else {
			stackTopTemp := temperatures[stack.topNode.val]
			if t <= stackTopTemp {
				stack.Push(i)
			} else {
				for t > stackTopTemp && stack.topNode.prev != nil {
					res[stack.topNode.val] = i - stack.topNode.val
					stack.Pop()
					stackTopTemp = temperatures[stack.topNode.val]
				}
				stack.Push(i)
			}
		}
	}

	for stack.topNode.prev != nil {
		res[stack.topNode.val] = 0
		stack.Pop()
	}

	return res
}
