package main

import (
	"fmt"
)

func main() {

	a := []int{2, 9, 2, 5, 6}

	/*
		a := [][]int{
			[]int{0, 1, 1, 1},
			[]int{1, 1, 1, 1},
			[]int{0, 1, 1, 1},
		}
	*/
	//words := []string{"a", "bb", "acda", "acae"}
	fmt.Printf("%v\n", minIncrementForUnique(a))
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
