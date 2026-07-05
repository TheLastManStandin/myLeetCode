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

	for left <= right {
		n1slice = (left + right) / 2
		n2slice = (nums1Len+nums2Len+1)/2 - n1slice

		Aleft := math.MinInt
		if n1slice > 0 {
			Aleft = nums1[n1slice-1]
		}

		Aright := math.MaxInt
		if n1slice < nums1Len {
			Aright = nums1[n1slice]
		}

		Bleft := math.MinInt
		if n2slice > 0 {
			Bleft = nums2[n2slice-1]
		}

		Bright := math.MaxInt
		if n2slice < nums2Len {
			Bright = nums2[n2slice]
		}

		if Aleft <= Bright && Bleft <= Aright {
			if (nums1Len+nums2Len)%2 == 0 {
				return (math.Max(float64(Aleft), float64(Bleft)) + math.Min(float64(Aright), float64(Bright))) / 2
			}
			return math.Max(float64(Aleft), float64(Bleft))
		} else if Aleft > Bright {
			right = n1slice
		} else if Bleft > Aright {
			left = n1slice + 1
		} else {
			break
		}
	}
	return 0
}
