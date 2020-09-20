package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	if m == 0 || nums1[m-1] == 0 {
		return getMedian(nums2)
	} else if len(nums2) == 0 {
		return getMedian(nums1)
	}

	low, hi, midSize, mid1, mid2 := 0, m, (m+n+1)>>1, m>>2, 0
	for mid1 <= hi && n-midSize >= 0 {
		mid2 = n - midSize
		if nums1[mid1] > nums2[mid2] {
			low = mid1
			mid1 = (low + hi + 1) >> 1
		}

	}

	orderArray := (append(left, right...))
	fmt.Println(orderArray)
	return getMedian(orderArray)
}

func getMedian(arr []int) float64 {
	length := len(arr)
	if length%2 == 0 {
		return float64((arr[length/2] + arr[length/2-1]) / 2)
	} else {
		return float64(arr[length/2])
	}
}

// func testFromMap(tests map[int][][]int, i int) {
// 	a, b := tests[i]
// 	return a, b
// }

func main() {
	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 4, 6, 8}
	// nums3 := []int{0}
	// nums5 := []int{0, 0}
	// nums6 := []int{}
	// nums7 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
