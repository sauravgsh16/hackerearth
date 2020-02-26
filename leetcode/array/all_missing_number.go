package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		fmt.Printf("%d  %d\n", nums[i], nums[nums[i]-1])

		for nums[i] != nums[nums[i]-1] {
			swap(nums, i, nums[i]-1)

			fmt.Printf("%v\n", nums)
		}
	}

	fmt.Printf("final : %v\n", nums)

	return []int{0}
}

func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
