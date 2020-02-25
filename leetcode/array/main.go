package main

func main() {
	//a := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//fmt.Printf("%d\n", minCostClimbingStairs(a))

	m := [][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}}

	//m := [][]int{[]int{1, 2}, []int{2, 1}}
	isToeplitzMatrix(m)
}

/*
arr[i] = 2
j = 1
arr[j] = 3
arr[j+1] = 2

j = 0
arr[i] = 2
arr[j] = 3
arr[j+1] = 3




3, 4, 2, 3
3, 2, 4, 3
2, 3, 4, 3
2, 3, 3, 4
             i	    e
[2, 6, 6, 4, 8, 10, 9, 12, 14, 15]
    s	     j


s=1
e=-1
    s
[1, 3, 2, 4, 5]

 s=0
 e=-1
[3, 2, 4, 5]

s=0
e=2
       e
[1, 2, 3, 5, 4, 4]


[T, T, T, F, T,  T, F,  T]

*/
