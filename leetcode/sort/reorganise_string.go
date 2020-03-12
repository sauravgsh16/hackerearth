package main

import (
	"strings"
)

func reorganizeString(s string) string {
	n := len(s)
	hash := make(map[rune]int)
	var max rune
	count := 0

	for _, r := range s {
		hash[r]++
		if hash[r] > count {
			count = hash[r]
			max = r
		}
	}

	res := make([]string, n)
	// maxCount := count[int(max)-'a']
	for i := 0; i < n; i += 2 {
		res[i] = string(max)
		count--
	}

	if count > 0 {
		return ""
	}

	delete(hash, max)

	j := 0
	for k := range hash {
		for hash[k] > 0 {
			if res[j] == "" {
				res[j] = string(k)
				hash[k]--
				j += 2
			} else {
				j++
			}

			if j > n-1 {
				j = 0
			}
		}
	}

	return strings.Join(res, "")
}
