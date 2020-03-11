class LinkedList:
    def __init__(self):
        self.head = None

class Node(object):
    def __init__(self, val, next=None, random=None):
        self.val = val
        self.next = next
        self.random = random


class Solution(object):
    def copyRandomList(self, head):
        random = {}

        curr = head

        while curr is not None:
            random[curr] = Node(curr.val)
            curr = curr.next

        print random

        curr = head
        while curr is not None:
            random.get(curr).next = random.get(curr.next)
            random.get(curr).random = random.get(curr.random, None)
            curr = curr.next
        
        return random.get(head)

def main():
    ll = LinkedList()

    l4 = Node(4)
    l3 = Node(3, next=l4)
    l2 = Node(2, next=l3)
    l1 = Node(1, next=l2)

    l1.random = l3
    l2.random = l1
    l3.random = None
    l4.random = l2

    ll.head = l1

    s = Solution()
    head = s.copyRandomList(ll.head)

    while head is not None:
        if head.random:
            print head.random.val,
        head = head.next


if __name__ == "__main__":
    main()
