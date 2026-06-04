package main

//func search(nums []int, target int, startsFromIndex ...int) int {
//	if startsFromIndex == nil {
//		startsFromIndex = make([]int, 1)
//		startsFromIndex[0] = 0
//	}
//
//	numsLen := len(nums)
//	if numsLen == 0 {
//		return -1
//	} else if numsLen == 1 {
//		if nums[0] == target {
//			return startsFromIndex[0]
//		}
//		return -1
//	}
//
//	midIndex := numsLen / 2
//	if nums[midIndex] == target {
//		return startsFromIndex[0] + midIndex
//	} else if nums[midIndex] < target {
//		return search(nums[midIndex:], target, startsFromIndex[0]+midIndex)
//	}
//	return search(nums[:midIndex], target, startsFromIndex[0])
//}

func searchMatrix(matrix [][]int, target int) bool {
	matrixLen := len(matrix)
	n := len(matrix[0])
	if matrixLen == 0 {
		return false
	} else if matrixLen == 1 {
		return search(matrix[0], target) != -1
	}

	midIndex := matrixLen / 2
	if matrix[midIndex][0] < target && target < matrix[midIndex][n-1] {
		return search(matrix[midIndex], target) != -1
	} else if matrix[midIndex][0] > target {
		return searchMatrix(matrix[:midIndex], target)
	}
	return searchMatrix(matrix[midIndex:], target)
}
