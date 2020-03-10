package main

import (
	"math"
)

// "krpgjbjjznpzdfy"
// "nxargkbydxmsgby"

// Proper solution
// For each iteration find the sum
func equalSubstrings(s, t string, maxCost int) int {
	n := len(s)
	var left, right int
	maxLenght := 0

	for right = 0; right < n; right++ {
		maxCost -= int(math.Abs(float64(int(s[right]) - int(t[right]))))
		if maxCost < 0 {
			maxCost += int(math.Abs(float64(int(s[left]) - int(t[left]))))
			left++
		}
		if maxLenght < (right - left + 1) {
			maxLenght = (right - left + 1)
		}
	}
	return maxLenght
}

func equalSubstringsMySolution(s, t string, maxCost int) int {
	maxlenght := 0
	n := len(s)
	i := 0
	count := 0
	cost := maxCost

	for i < n {
		val := int(math.Abs(float64(int(s[i]) - int(t[i]))))

		if cost-val >= 0 {
			cost -= val
			count++
		} else {
			if maxlenght < count {
				maxlenght = count
			}

			if val > maxCost {
				cost = maxCost
				count = 0
			} else if val == maxCost {
				cost = maxCost
				count = 1
			} else {
				count = 0
				temp := maxCost
				j := i
				for j > 0 && temp-val >= 0 {
					temp -= val
					count++
					j--
					val = int(math.Abs(float64(int(s[j]) - int(t[j]))))
				}
				cost = temp
			}
		}
		i++
	}

	if maxlenght < count {
		maxlenght = count
	}

	return maxlenght
}

//  cd
// bcdf
