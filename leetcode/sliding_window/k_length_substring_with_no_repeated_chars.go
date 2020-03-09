package main

// geeksforgeeks
func kLengthSubstringNoRepeatedChar(s string) int {
	n := len(s)
	curlen := 1
	maxLen := 1
	prevIndex := 0

	visited := make([]int, 256)

	for i := range visited {
		visited[i] = -1
	}

	// mark the first char in visited as 0
	visited[s[0]] = 0

	for i := 1; i < n; i++ {
		prevIndex = visited[s[i]]

		// if char is not present in visited or
		// if char is not part of the substring being checked
		if prevIndex == -1 || (i-curlen) > prevIndex {
			curlen++
		} else {
			if curlen > maxLen {
				maxLen = curlen
			}

			// update curlen as 1
			curlen = 1
		}

		// update visited with index of char
		visited[s[i]] = i
	}
	return maxLen
}
