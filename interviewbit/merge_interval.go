package main

func insert(intervals [][]int, in []int) [][]int {

	new := []int{0, 0}

	begin := [][]int{}
	end := [][]int{}

	for _, curr := range intervals {
		if in[0] > curr[0] && in[1] < curr[1] {
			return intervals
		}

		// if start of range > end of current range
		if in[0] > curr[1] {
			begin = append(begin, curr)

			// if end if less than current range
		} else if in[1] < curr[0] {
			end = append(end, curr)

		} else if in[0] > curr[0] && in[0] < curr[1] {
			new[0] = curr[0]
		} else if in[1] > curr[0] && in[1] < curr[1] {
			new[1] = curr[1]
		}
	}

	if new[0] == 0 {
		new[0] = in[0]
	}

	if new[1] == 0 {
		new[1] = in[1]
	}

	result := append(append(begin, new), end...)

	return result
}
