package main

func missingNumber(nums []int) int {
	max := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		sum += nums[i]
	}

	total := (max * (max + 1)) / 2

	return total - sum
}
