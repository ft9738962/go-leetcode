package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	if m == 0 {
		return getMedian(nums2)
	} else if nums1[m-1] == 0 {
		return getMedian(nums2)
	} else if len(nums2) == 0 {
		return getMedian(nums1)
	}

	low, hi, midSize, mid1 := 0, m-1, (m+n+1)>>1, m>>1
	mid2 := midSize - mid1
	for nums1[mid1-1] > nums2[mid2] || nums1[mid1] < nums2[mid2-1] {
		if nums1[mid1-1] > nums2[mid2] {
			hi = hi >> 1
		} else if nums1[mid1] < nums2[mid2-1] {
			low = mid1
		}
		mid1 = (low + hi + 1) >> 1
		if mid1 == 0 {
			return getMedian(append(nums2, nums1...))
		} else if mid1 == hi {
			return getMedian(append(nums1, nums2...))
		}
		mid2 = midSize - mid1
		fmt.Println(mid1, mid2)
	}
	return getMedian2(nums1[mid1-1], nums1[mid1],
		nums2[mid2-1], nums2[mid2], n+m)
}

func getMedian(arr []int) float64 {
	length := len(arr)
	if length%2 == 0 {
		return float64((arr[length/2] + arr[length/2-1]) / 2)
	}
	return float64(arr[length/2])
}

func getMedian2(nums1Left, nums1Right, nums2Left, nums2Right, totLength int) float64 {
	mi := math.Max(float64(nums1Left), float64(nums2Left))
	if totLength%2 == 1 {
		return mi
	}
	ma := math.Min(float64(nums1Right), float64(nums2Right))
	return (mi + ma) / 2
}

func main() {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4}
	// nums3 := []int{0}
	// nums5 := []int{0, 0}
	// nums6 := []int{}
	// nums7 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
