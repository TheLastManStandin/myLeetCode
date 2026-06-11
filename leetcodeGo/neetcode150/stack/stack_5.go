package stack

type stack []float32

func (s *stack) Push(v float32) {
	*s = append(*s, v)
}

func (s *stack) Pop() float32 {
	lens := len(*s)
	res := (*s)[lens-1]
	*s = (*s)[:lens-1]
	return res
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	res := 0
	m := make([]float32, target)
	s := stack{}

	for i := n - 1; i >= 0; i-- {
		carPos := position[i]
		carSpeed := speed[i]

		timeToRichTarget := float32(target-carPos) / float32(carSpeed)

		m[carPos] = timeToRichTarget
	}

	for i := target - 1; i >= 0; i-- {
		if m[i] == 0 {
			continue
		}
		if len(s) > 0 {
			stackTop := s.Pop()
			if stackTop < m[i] {
				res++
				s.Push(m[i])
			} else {
				s.Push(stackTop)
			}
		} else {
			s.Push(m[i])
		}
	}

	if len(s) > 0 {
		res++
	}

	return res
}
