package main

import (
	"fmt"
)

func main() {
	/*
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')

		txt, _ := reader.ReadString('\n')

		arr := []int{}
		for _, f := range strings.Fields(txt) {
			i, _ := strconv.Atoi(string(f))
			arr = append(arr, i)
		}
	*/

	arr := []int{1, 2, 3, 4, 5, 6, 2, 9, 7, 8}

	findNextGreaterSum(arr)
}

type stack struct {
	arr []int
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
}

func (s *stack) pop() int {
	e := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return e
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func findNextGreaterSum(arr []int) {
	n := len(arr)
	allx := []int{}
	ally := []int{}

	r := &stack{}
	r.push(0)

	// no values before it
	allx = append(allx, -1)

	for i := 1; i < n; i++ {
		next := arr[i]
		var x int

		if !r.isEmpty() {

			e := r.pop()

			for arr[e] < next {
				if r.isEmpty() {
					break
				}
				e = r.pop()
			}

			if arr[e] > next {
				r.push(e)
				x = e + 1
			} else {
				x = -1
			}
		}

		allx = append(allx, x)

		r.push(i)
	}

	l := &stack{}
	l.push(n - 1)
	//ally = append(ally, -1)

	for i := n - 1; i >= 0; i-- {
		next := arr[i]
		var y int

		if !l.isEmpty() {

			e := l.pop()

			for arr[e] < next {
				if l.isEmpty() {
					break
				}
				e = l.pop()
			}

			if arr[e] > next {
				l.push(e)
				y = e + 1
			} else {
				y = -1
			}
		}

		ally = append(ally, y)
		l.push(i)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", allx[i]+ally[n-1-i])
	}

	fmt.Println()
}
