package main

import "fmt"

func main() {
	/*
			l7 := &ListNode{
				Val: 8,
				// Next: l2,
			}

			l6 := &ListNode{
				Val:  7,
				Next: l7,
			}

			l5 := &ListNode{
				Val:  6,
				Next: l6,
			}

			l4 := &ListNode{
				Val:  5,
				Next: l5,
			}


		l3 := &ListNode{
			Val: 4,
			Next: l4,
		}
	*/

	l2 := &ListNode{
		Val: 3,
		// Next: l3,
	}
	l1 := &ListNode{
		Val:  2,
		Next: l2,
	}

	h := &ListNode{
		Val:  1,
		Next: l1,
	}
	//fmt.Printf("%v\n")
	head := reverseListIterative(h)

	for head != nil {
		fmt.Printf("%d\n", head.Val)
		head = head.Next
	}

}

// ListNode of a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
