package main

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	stackLeft := make([]int, 0, len(height))
	stackRight := make([]int, 0, len(height))
	ans := 0

	if len(height) < 3 {
		return ans
	}

	stackRight = append(stackRight, height[right])
	right--
	stackLeft = append(stackLeft, height[left])
	left++

	for left <= right {
		if stackRight[0] < stackLeft[0] {
			if height[right] > stackRight[0] {
				stackRightLen := len(stackRight)
				for i := 1; i <= stackRightLen-1; i++ {
					ans += stackRight[0] - stackRight[i]
				}
				stackRight = make([]int, 0, len(height))
				stackRight = append(stackRight, height[right])
			} else {
				stackRight = append(stackRight, height[right])
			}
			right--
		} else {
			if height[left] > stackLeft[0] {
				stackLeftLen := len(stackLeft)
				for i := 1; i <= stackLeftLen-1; i++ {
					ans += stackLeft[0] - stackLeft[i]
				}
				stackLeft = make([]int, 0, len(height))
				stackLeft = append(stackLeft, height[left])
			} else {
				stackLeft = append(stackLeft, height[left])
			}
			left++
		}
	}

	curLevel := min(stackRight[0], stackLeft[0])
	if left == right && height[left] <= curLevel {
		stackRight = append(stackRight, height[left])
	}

	stackRightLen := len(stackRight)
	for i := 1; i <= stackRightLen-1; i++ {
		if curLevel-stackRight[i] >= 0 {
			ans += curLevel - stackRight[i]
		}
	}

	stackLeftLen := len(stackLeft)
	for i := 1; i <= stackLeftLen-1; i++ {
		if curLevel-stackLeft[i] >= 0 {
			ans += curLevel - stackLeft[i]
		}
	}

	return ans
}
