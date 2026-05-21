package main

import (
	"sort"
)

func findSliderValInSlice(sliderVal int, numsSlice []int) [][]int {
	ans := make([][]int, 0)
	leftPointer := 0
	rightPointer := len(numsSlice) - 1

	for leftPointer < rightPointer {
		summ := numsSlice[leftPointer] + numsSlice[rightPointer]

		if summ > -sliderVal {
			rpVal := numsSlice[rightPointer]
			for numsSlice[rightPointer] == rpVal && rightPointer > 0 {
				rightPointer--
			}
		} else if summ < -sliderVal {
			lpVal := numsSlice[leftPointer]
			for numsSlice[leftPointer] == lpVal && leftPointer < len(numsSlice) {
				leftPointer++
			}
		} else {
			ans = append(ans, []int{sliderVal, numsSlice[leftPointer], numsSlice[rightPointer]})
			rpVal := numsSlice[rightPointer]
			for numsSlice[rightPointer] == rpVal && rightPointer > 0 {
				rightPointer--
			}
		}
	}

	return ans
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([][]int, 0)

	slider := 0

	for sliderVal := nums[slider]; sliderVal <= 0; {
		if len(nums[slider:]) < 3 {
			break
		}

		ans = append(ans, findSliderValInSlice(sliderVal, nums[slider+1:])...)

		for nums[slider] == sliderVal && slider < len(nums)-1 {
			slider++
		}
		sliderVal = nums[slider]
	}

	return ans
}
