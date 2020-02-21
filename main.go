package main

import (
	"fmt"
)

type stack struct {
	d []int
}

func (s *stack) push(i int) {
	s.d = append(s.d, i)
}

func (s *stack) pop() int {
	i := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return i
}

func (s *stack) isEmpty() bool {
	return len(s.d) == 0
}

func (s *stack) top() int {
	return s.d[len(s.d)-1]
}

func nextElement(arr []int, order string) []int {
	s := &stack{}
	ans := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {

		if order == "G" {

			for !s.isEmpty() && arr[s.top()] <= arr[i] {
				s.pop()
			}

		} else {

			for !s.isEmpty() && arr[s.top()] >= arr[i] {
				s.pop()
			}

		}

		if s.isEmpty() {
			ans[i] = -1
		} else {
			ans[i] = s.top()
		}

		// push index
		s.push(i)
	}

	fmt.Println(ans)

	return ans
}

func main() {
	var (
		loop int
		d    int
	)

	fmt.Scanf("%d", &loop)

	arr := make([]int, 0)
	for i := 0; i < loop; i++ {
		fmt.Scanf("%d", &d)
		arr = append(arr, d)
	}

	greater := nextElement(arr, "G")
	smaller := nextElement(arr, "S")

	for i := 0; i < len(arr); i++ {
		if greater[i] != -1 && smaller[greater[i]] != -1 {
			fmt.Printf("%d ", arr[smaller[greater[i]]])
		} else {
			fmt.Printf("%d ", -1)
		}
	}
	fmt.Println()
}
