package main

import "fmt"

func main() {
	/*
		a1 := []int{1, 2, 3, 0, 0, 0, 0}
		a2 := []int{2, 4, 6, 7}
	*/
	fmt.Printf("%d\n", missingNumber([]int{2}))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if _, ok := hash[nums[i]]; ok {
			if i-hash[nums[i]] == k {
				fmt.Printf("%v\n", hash)

				return true
			}
		}
		hash[nums[i]] = i
	}

	fmt.Printf("%v -- returning false\n", hash)
	return false
}

func plusOne(arr []int) []int {
	n := len(arr)
	inc := 1

	for i := n - 1; i >= 0; i-- {
		arr[i] += inc
		if arr[i] >= 10 {
			arr[i] -= 10
			continue
		}
		inc = 0
		break
	}

	var res []int

	if inc == 1 {
		res = []int{1}
		for _, i := range arr {
			res = append(res, i)
		}
	} else {
		res = arr
	}
	return res
}
