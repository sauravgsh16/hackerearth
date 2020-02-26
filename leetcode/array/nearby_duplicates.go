package main

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if _, ok := hash[nums[i]]; ok {
			if i-hash[nums[i]] == k {
				return true
			}
		}
		hash[nums[i]] = i
	}
	return false
}
