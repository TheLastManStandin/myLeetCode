package main

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	stackLeft := make([]int, 0, len(height))
	stackRight := make([]int, 0, len(height))
	ans := 0

	for left < right {
		if len(stackLeft) == 0 {
			stackLeft = append(stackLeft, height[left])
		} else {
			if height[left] > stackLeft[0] {
				stackLeftLen := len(stackLeft)
				for i := 1; i < stackLeftLen-1; i++ {
					ans += stackLeft[0] - stackLeft[i]
				}
				stackLeft = make([]int, 0, len(height))
				stackLeft = append(stackLeft, height[left])
			} else {
				stackLeft = append(stackLeft, height[left])
			}
		}

		if len(stackRight) == 0 {
			stackRight = append(stackRight, height[right])
		} else {
			if height[right] > stackRight[0] {
				stackRightLen := len(stackRight)
				for i := 1; i < stackRightLen-1; i++ {
					ans += stackRight[0] - stackRight[i]
				}
				stackRight = make([]int, 0, len(height))
				stackRight = append(stackRight, height[right])
			} else {
				stackRight = append(stackRight, height[right])
			}
		}

		left++
		right--
	}

	if left == right {
		stackRight = append(stackRight, height[left])
	}
	curLevel := min(stackRight[0], stackLeft[0])

	stackRightLen := len(stackRight)
	for i := 1; i < stackRightLen-1; i++ {
		ans += curLevel - stackRight[i]
	}

	stackLeftLen := len(stackLeft)
	for i := 1; i < stackLeftLen-1; i++ {
		ans += stackLeft[0] - stackLeft[i]
	}

	return ans
}
