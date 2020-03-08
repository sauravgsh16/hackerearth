package main

import (
	"fmt"
)

func main() {

	a := []int{2, 0, 1}
	fmt.Printf("%v\n", sortBalls(a))
}

func matchSubsequence(s string, words []string) int {
	wordCount := make(map[string][]string)

	for _, word := range words {
		key := string(word[0])
		if _, ok := wordCount[key]; !ok {
			wordCount[key] = []string{}
		}
		wordCount[key] = append(wordCount[key], word[1:])
	}

	count := 0
	for _, ch := range s {
		strCh := string(ch)

		toCheck := wordCount[strCh]
		delete(wordCount, strCh)

		for _, word := range toCheck {
			if len(word) == 0 {
				count++
			} else {
				key := string(word[0])
				wordCount[key] = append(wordCount[key], word[1:])
			}
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
