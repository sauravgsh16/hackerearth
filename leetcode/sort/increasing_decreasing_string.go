package main

func sortString(s string) string {
	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		if count[int(val)] == 0 {
			count[int(val)] = 1
		} else {
			count[int(val)]++
		}
	}

	res := ""
	i := 0
	for i < len(s) {
		for j := 0; j < len(count); j++ {
			if count[j] > 0 {
				res += string(j + 'a')
				count[j]--
				i++
			}
		}

		for j := len(count) - 1; j >= 0; j-- {
			if count[j] > 0 {
				res += string(j + 'a')
				count[j]--
				i++
			}
		}
	}

	return res
}
