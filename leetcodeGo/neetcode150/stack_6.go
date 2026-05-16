package main

//type doubleStack [][2]int
//
//func (s *doubleStack) Push(v [2]int) {
//	*s = append(*s, v)
//}
//
//func (s *doubleStack) Pop() [2]int {
//	lens := len(*s)
//	res := (*s)[lens-1]
//	*s = (*s)[:lens-1]
//	return res
//}

type Bar struct {
	height int
	index  int
}
type doubleLifo []Bar

func (dl *doubleLifo) GetOut() Bar {
	return (*dl)[0]
}

func (dl *doubleLifo) Pop() {
	*dl = (*dl)[1:]
}

func (dl *doubleLifo) Push(x Bar) {
	*dl = append(*dl, x)
}

func largestRectangleArea(heights []int) int {
	dl := doubleLifo{}
	ans := 0
	squares := make([]int, len(heights))

	for i, h := range heights {
		if len(dl) == 0 {
			dl.Push(Bar{height: h, index: i})
		} else {
			dlVal := dl.GetOut()
			if dlVal.height <= h {
				dl.Push(Bar{height: h, index: i})
			} else {
				for dlVal.height > h {
					dl.Pop()
					squares[dlVal.index] = (i - dlVal.index) * dlVal.height
					dlVal = dl.GetOut()
				}
			}
		}
	}

	return ans
}
