package main

func subarraysDivByK(arr []int, K int) {
	sum, count := 0, 0

	hash := make(map[int]int)
	hash[0] = 1

	for _, a := range arr {
		sum = (sum + a) % K
		if sum < 0 {
			sum += K
		}
		count += hash[sum]
		hash[sum]++
	}
}
