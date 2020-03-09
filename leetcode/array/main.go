package main

func main() {

	//a := []int{10, 1, 7}
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

func grumpyOwner(cust []int, grumpy []int, x int) int {
	n := len(cust)
	count := 0

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			count += cust[i]
			cust[i] = 0
		}
	}

	var max, i int
	for i = 0; i < x; i++ {
		max += cust[i]
	}
	sum := max

	for j := i; j < n; j++ {
		sum = sum + cust[j] - cust[j-x]
		if sum > max {
			max = sum
		}
	}

	return count + max
}
