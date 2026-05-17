package main

import "slices"

type Bar struct {
	height int
	index  int
}

type barStack []Bar

func (bs *barStack) getOut() Bar {
	return (*bs)[len(*bs)-1]
}

func (bs *barStack) push(b Bar) {
	*bs = append(*bs, b)
}

func (bs *barStack) pop() {
	*bs = (*bs)[:len(*bs)-1]
}

func largestRectangleArea(heights []int) int {
	bs := new(barStack)
	squares := make([]int, len(heights))

	// Проходимся слева направо
	for i, h := range heights {
		if len(*bs) == 0 {
			bs.push(Bar{height: h, index: i})
		} else {
			dlVal := bs.getOut()
			if dlVal.height <= h {
				bs.push(Bar{height: h, index: i})
			} else {
				for dlVal.height > h {
					bs.pop()
					squares[dlVal.index] = (i - dlVal.index) * dlVal.height
					if len(*bs) > 0 {
						dlVal = bs.getOut()
					} else {
						dlVal.height = -1
					}
				}
				bs.push(Bar{height: h, index: i})
			}
		}
	}

	dlVal := bs.getOut()
	bsLen := len(*bs)
	for i := 0; i < bsLen; i++ {
		bs.pop()
		squares[dlVal.index] += (len(heights) - dlVal.index) * dlVal.height
		if len(*bs) > 0 {
			dlVal = bs.getOut()
		} else {
			dlVal.height = -1
		}
	}

	bs = new(barStack)

	// Проходимся справа налево
	for i := len(heights) - 1; i >= 0; i-- {
		h := heights[i]
		if len(*bs) == 0 {
			bs.push(Bar{height: h, index: i})
		} else {
			dlVal = bs.getOut()
			if dlVal.height <= h {
				bs.push(Bar{height: h, index: i})
			} else {
				for dlVal.height > h {
					bs.pop()
					squares[dlVal.index] += (dlVal.index - i - 1) * dlVal.height
					if len(*bs) > 0 {
						dlVal = bs.getOut()
					} else {
						dlVal.height = -1
					}
				}
				bs.push(Bar{height: h, index: i})
			}
		}
	}

	dlVal = bs.getOut()
	bsLen = len(*bs)
	for i := 0; i < bsLen; i++ {
		bs.pop()
		squares[dlVal.index] += dlVal.index * dlVal.height
		if len(*bs) > 0 {
			dlVal = bs.getOut()
		} else {
			dlVal.height = -1
		}
	}

	return slices.Max(squares)
}
