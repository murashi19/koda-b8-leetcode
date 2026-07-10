package main

import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := slices.Concat(nums1, nums2)
	slices.Sort(merge)

	n := len(merge)
	mid := n / 2

	var med float64
	if n%2 == 0 {
		med = (float64(merge[mid-1]+merge[mid]) / float64(2))
	} else {
		med = (float64(merge[mid]))
	}
	return med
}
