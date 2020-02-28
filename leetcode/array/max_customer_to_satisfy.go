package main

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)

	satisfied := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
			customers[i] = 0
		}
	}

	rollingSum := 0
	maxCanSatisfy := 0
	for i := 0; i < n; i++ {
		rollingSum += customers[i]

		if i > X-1 {
			rollingSum -= customers[i-X]
		}
		if rollingSum > maxCanSatisfy {
			maxCanSatisfy = rollingSum
		}
	}

	return satisfied + maxCanSatisfy
}
