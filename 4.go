package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	if m == 0 || nums1[m] == 0 {
		if n%2 == 0 {
			return float64(nums2[n/2]+nums2[n/2-1]) / 2
		} else {
			return float64(nums2[n/2])
		}
	} else if len(nums2) == 0 {
		if n%2 == 0 {
			return float64(nums1[n/2]+nums1[n/2-1]) / 2
		} else {
			return float64(nums1[n/2])
		}
	}

}
