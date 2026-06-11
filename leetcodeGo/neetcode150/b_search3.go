package main

import "slices"

func b_search(piles []int, h int) int {
	var left, lowest int
	right := slices.Max(piles) + 1

	for left < right {
		mid := left + (right-left)/2
		if mid == 0 {
			break
		}
		calcedH := calcWithK(piles, mid)
		if calcedH > h {
			left = mid + 1
		} else {
			lowest = mid
			right = mid
		}
	}

	if lowest != 0 {
		return lowest
	}

	return -1
}

func calcWithK(piles []int, k int) int {
	res := 0
	for _, v := range piles {
		res += (v + k - 1) / k
	}
	return res
}

func minEatingSpeed(piles []int, h int) int {
	return b_search(piles, h)
}
