package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var prev, next, indexNums1, indexNums2 int
	length := len(nums1) + len(nums2)
	median := length / 2

	for i := 0; i <= median; i++ {
		prev = next

		if indexNums1 < len(nums1) && indexNums2 < len(nums2) {
			if nums1[indexNums1] < nums2[indexNums2] {
				next = nums1[indexNums1]
				indexNums1++
			} else {
				next = nums2[indexNums2]
				indexNums2++
			}
		} else if indexNums1 < len(nums1) {
			next = nums1[indexNums1]
			indexNums1++
		} else if indexNums2 < len(nums2) {
			next = nums2[indexNums2]
			indexNums2++
		}
	}

	if length%2 == 0 {
		return float64(prev+next) / 2.0
	}

	return float64(next)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
