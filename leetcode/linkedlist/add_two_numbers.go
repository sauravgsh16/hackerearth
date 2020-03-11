package main

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	var carry int

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		curr.Next = &ListNode{
			Val: sum,
		}

		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}

func addTwoNumbersMySolution(l1, l2 *ListNode) *ListNode {
	var sum, carry int
	head := &ListNode{}

	curr := head

	for l1 != nil && l2 != nil {
		if carry > 0 {
			sum = carry + l1.Val + l2.Val
		} else {
			sum = l1.Val + l2.Val
		}
		unit := sum % 10
		carry = sum / 10

		curr.Val = unit

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}

	if carry > 0 && (l1 != nil || l2 != nil) {
		for l1 != nil {
			sum = carry + l1.Val
			unit := sum % 10
			carry = sum / 10

			curr.Val = unit
			l1 = l1.Next
			if l1 != nil {
				curr.Next = &ListNode{}
				curr = curr.Next
			}
		}

		for l2 != nil {
			sum = carry + l2.Val
			unit := sum % 10
			carry = sum / 10

			curr.Val = unit
			l2 = l2.Next
			if l2 != nil {
				curr.Next = &ListNode{}
				curr = curr.Next
			}
		}
	} else if l1 != nil {
		*curr = *l1
	} else if l2 != nil {
		*curr = *l2
	}

	if l1 == nil && l2 == nil && carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
