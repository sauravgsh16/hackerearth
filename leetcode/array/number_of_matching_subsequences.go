package main

func numMatchingSubseq(s string, words []string) int {
	wordChecker := make(map[string][]string)

	// Prepare map , with key as the first letter, and value as array of remaining letters
	for _, word := range words {
		key := string(word[0])
		if _, ok := wordChecker[key]; !ok {
			wordChecker[key] = []string{}
		}
		wordChecker[key] = append(wordChecker[key], word[1:])
	}

	count := 0
	for _, ch := range s {
		strCh := string(ch)

		// Get the array, which the character needs to be search
		toCheck := wordChecker[strCh]
		// Delete array from map, so that, if the same character
		// appears again, we do not append to same array
		delete(wordChecker, strCh)

		for _, word := range toCheck {
			if len(word) == 0 {
				count++
			} else {
				key := string(word[0])
				// create new array for the remaining characters to be checked
				wordChecker[key] = append(wordChecker[key], word[1:])
			}
		}
	}

	return count
}
