package b_search

func searchShift(nums []int) int {
	left := 0
	right := len(nums)
	lastMid := left

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[lastMid] {
			left = mid + 1
		} else if nums[mid] <= nums[lastMid] {
			lastMid = mid
			right = mid
		}
	}
	return lastMid
}

func search(nums []int, target int) int {
	shift := searchShift(nums)
	newNums := make([]int, 0, len(nums))
	newNums = append(newNums, append(nums[shift:], nums[:shift]...)...)

	left := 0
	right := len(nums)
	ans := -1

	for left < right {
		mid := left + (right-left)/2
		midVal := newNums[mid]

		if midVal == target {
			ans = mid
			break
		} else if midVal > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if ans != -1 {
		negativeShift := len(nums) - shift
		if ans < negativeShift {
			return ans + shift
		}
		return ans - negativeShift
	}
	return ans
}
