package main

func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	maxS := 0

	for left < right {
		s := min(heights[left], heights[right]) * (right - left)
		maxS = max(maxS, s)
		if heights[left] < heights[right] {
			left++
		} else if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}

	return maxS
}
