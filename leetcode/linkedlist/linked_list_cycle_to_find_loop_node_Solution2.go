package main

/*
Algorithm:

   <------m----><--k-->
   1 --> 2 --> 3 --> 4 <--(Slow/fast)
               ^
               |     |
               |     |
	       6 <-- 5
   n = lenght of cycle

Distance traveled by fast pointer = 2 * (Distance traveled
                                         by slow pointer)

(m + n*x + k) = 2*(m + n*y + k)

Note that before meeting the point shown above, fast
was moving at twice speed.

x -->  Number of complete cyclic rounds made by
       fast pointer before they meet first time

y -->  Number of complete cyclic rounds made by
       slow pointer before they meet first time

From above equation, we can conclude below

    m + k = (x-2y)*n

Which means m+k is a multiple of n.

So if we start moving both pointers again at same speed such that one pointer (say slow)
begins from head node of linked list and other pointer (say fast) begins from meeting point.
When slow pointer reaches beginning of loop (has made m steps), fast pointer would have made
also moved m steps as they are now moving same pace. Since m+k is a multiple of n and fast
starts from k, they would meet at the beginning. Can they meet before also?
No because slow pointer enters the cycle first time after m steps.

*/

func detectCycleOptimised(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if slow == fast {
		slow = head

		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}

		return fast
	}
	return nil
}
