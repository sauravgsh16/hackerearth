package main

func majorityElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	hash := make(map[int]int)

	ele := nums[0]
	max := 1

	for i := 0; i < n; i++ {
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = 1
		} else {
			hash[nums[i]]++
			if hash[nums[i]] > max {
				max = hash[nums[i]]
				ele = nums[i]
			}
		}
	}

	return ele
}
