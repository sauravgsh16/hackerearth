package main

/*
Given a string s that consists of only uppercase English letters, you can perform at most k operations on that string.

In one operation, you can choose any character of the string and change it to any other uppercase English character.

Find the length of the longest sub-string containing all repeating letters you can get after performing the above operations.

Note:
Both the string's length and k will not exceed 104.

Example 1:

Input:
s = "ABAB", k = 2

Output:
4

Explanation:
Replace the two 'A's with two 'B's or vice versa.

*/

func characterReplacement(s string, k int) int {
	var maxCount int

	n := len(s)
	count := make([]int, 26)
	left, right := 0, 0

	// The idea is the find a window, which satisfies the condition
	// of taking the maxCount + 'k', it should be less than or equal to the
	// difference between the right and left pointer
	// If (maxCount + k) becomes less than the length of subarray,
	// we decrement the value of char pointed to be the left, and increment left pointer

	for right = 0; right < n; right++ {
		idx := s[right] - 'A'
		count[idx]++

		maxCount = max(maxCount, count[idx])

		if (maxCount + k) <= (right - left) {
			idx = s[left] - 'A'
			count[idx]--
			left++
		}
	}

	return n - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
