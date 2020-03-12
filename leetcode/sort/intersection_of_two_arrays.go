package main

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:

    Each element in the result must be unique.
    The result can be in any order.


*/

import "sort"

func intersection(arr1, arr2 []int) []int {
	res := []int{}

	sort.Ints(arr1)
	sort.Ints(arr2)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			if len(res) == 0 || res[len(res)-1] != arr1[i] {
				res = append(res, arr1[i])
			}
			i++
			j++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			i++
		}
	}

	return res
}
