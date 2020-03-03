package main

// Similar to buy and sell stock, but instead of min price, we need to keep track of max value.
// When keeping track of max value, we need to take care decay effect every step due to the distance property
func maxScoreSightseeingPair(arr []int) int {
	maxIndex := arr[0] - 1
	var maxScore int

	for j := 1; j < len(arr); j++ {

		score := maxIndex + arr[j]
		if score > maxScore {
			maxScore = score
		}

		if maxIndex < arr[j] {
			maxIndex = arr[j]
		}

		maxIndex-- // Index penalty as we move each step
	}
	return maxScore
}
