package b_search

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	if nums1Len > nums2Len {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	}

	n1slice, n2slice := 0, 0 // до i элемента
	left, right := 0, nums1Len
	Aleft := -99999999
	Aright := 99999999

	for left < right {
		n1slice = (left + right) / 2
		n2slice = (nums1Len+nums2Len+1)/2 - n1slice

		Aleft = nums1[n1slice]
		Aright = nums1[n1slice+1]
		Bleft := nums2[n2slice]
		Bright := nums2[n2slice+1]

		if Aleft <= Bright && Bleft <= Aright {
			break
		} else if Aleft > Bright {
			right = n1slice
		} else if Bleft > Aright {
			left = n1slice + 1
		} else {
			break
		}
	}
}
