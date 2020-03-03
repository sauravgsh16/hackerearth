package main

// Requirement:
//      1) Constant Space
//	2) Array should not be modified(read only)
//	3) runtime complexity should be less than O(n^2)

// Method 1:
// Using the pigeon hole principle, which states that:
// If there are n items and m holes such that n>m, then
// at least one hole must contain multiple objects
func findDuplicatePigeonHole(nums []int) int {
	n := len(nums)

	low := 0
	high := n - 1

	for low < high {
		mid := low + (high-low)/2

		var count int
		for _, i := range nums {
			if i <= mid {
				count++
			}
		}
		if count <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

// Method 2: TORTOISE and the HARE **********
func findDuplicateFastSlowPointer(nums []int) int {
	n := len(nums)

	// a := []int{1, 5, 3, 2, 4, 6, 7, 3}
	if n > 1 {
		// The "tortoise and hare" step. We start at the end of the array
		// and try to find an intersection point
		slow := nums[0]
		fast := nums[nums[0]]

		// Keep advancing the 'slow' by one step and the 'fast' by two steps
		// until they meet inside the loop
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
		}

		// Start another pointer from end of array and move it forward until
		// it hits the pointer inside the array
		finder := 0
		for slow != finder {
			slow = nums[slow]
			finder = nums[finder]
		}
		return slow

	}
	return -1
}
