package main

import (
	"math"
)

// "krpgjbjjznpzdfy"
// "nxargkbydxmsgby"

func equalSubstrings(s, t string, maxCost int) int {
	maxlenght := 0
	n := len(s)
	i := 0
	count := 0
	cost := maxCost

	for i < n {
		val := int(math.Abs(float64(int(s[i]) - int(t[i]))))

		if cost >= val {
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
				for j > 0 && temp >= 0 {
					j--
					val = int(math.Abs(float64(int(s[j]) - int(t[j]))))
					if temp-val < 0 {
						break
					}
					temp -= val
					count++
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
