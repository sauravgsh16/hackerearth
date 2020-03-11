package main

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

	l12 := &ListNode{
		Val: 8,
		//Next: l3,
	}
	l11 := &ListNode{
		Val:  2,
		Next: l12,
	}

	h1 := &ListNode{
		Val:  4,
		Next: l11,
	}

	/*
				l22 := &ListNode{
					Val: 3,
				}

				l21 := &ListNode{
					Val:  4,
					Next: l22,
				}


			h2 := &ListNode{
				Val: 2,
				//Next: l21,
			}
			//fmt.Printf("%v\n")
			head := addTwoNumbers(h1, h2)
		for head != nil {
			fmt.Printf("%d\n", head.Val)
			head = head.Next
		}
	*/

	//t := createTree()
	s := []int{1, 2, 4, 2, 8, 4, 2, 4}
	//fmt.Printf("%t\n", )
	kmpSearchString(h1, s)

}

// ListNode of a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
