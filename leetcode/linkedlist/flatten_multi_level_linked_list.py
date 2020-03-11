class LinkedList:
    def __init__(self):
        self.head = None


class ListNode:
    def __init__(self, val):
        self.val = val
        self.prev = None
        self.next = None
        self.child = None

class Solution(object):
    def flatten(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if head is None:
            return
        self.flattentail(head)
        return head
    
    def flattentail(self, head):
        if head.next is None:
            return head
    
        curr = head

        child, prev = None, curr.next

        if curr.child:
            child = self.flattentail(curr.child)
            curr.child = None

        while child is not None:
            curr.next = child
            child = child.next
            curr = curr.next

        curr.next = self.flattentail(prev)

        return head
    

def main():
    ll = LinkedList()
    ll.head = ListNode(1)
    ll.head.next = ListNode(2)
    ll.head.next.next = ListNode(3)
    ll.head.next.next.child = ListNode(4)
    ll.head.next.next.child.next = ListNode(5)
    ll.head.next.next.next = ListNode(6)

    s = Solution()
    head = s.flatten(ll.head)

    while head is not None:
        print head.val,
        head = head.next

if __name__ == "__main__":
    main()