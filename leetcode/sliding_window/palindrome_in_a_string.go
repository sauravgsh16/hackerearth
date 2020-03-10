package main

func ckeckIfPalindrome(s1, s2 string) bool {
	count := make([]int, 26)
	current := make([]int, 26)
	n := len(s2)

	for _, r := range s1 {
		count[r-'a']++
	}

	for i := 0; i < n; i++ {
		current[s2[i]-'a']++

		// check if passed window
		if (i - len(s1)) >= 0 {
			// decrement the count of i-len(s1) index from current
			current[s2[i-len(s1)]-'a']--
		}

		if isEqual(current, count) {
			return true
		}
	}

	return false
}

func isEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
