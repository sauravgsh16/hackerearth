package main

func leastInterval(tasks []byte, n int) int {
	count := make(map[byte]int, 0)

	for _, t := range tasks {
		if _, ok := count[t]; !ok {
			count[t] = 0
		}
		count[t]++
	}

	max := len(count) * nl
	result := make([]byte, max)

	for k, v := range count {
		for i := 0; i < v; i++ {

		}
	}
}
