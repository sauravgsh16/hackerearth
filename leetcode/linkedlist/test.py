def isSubPath(head):
    A, dp = [head.val], [0]
    i = 0
    node = head.next
    while node:
        while node and i and node.val != A[i]:
            i = dp[i - 1]
        i += node.val == A[i]
        A.append(node.val)
        dp.append(i)
        node = node.next
    print dp


class Node:
    def __init__(self, val):
        self.val = val
        self.next = None


class LL:
    def __init__(self):
        self.head = None

    def insert(self, val):
        n = Node(val)
        if self.head is None:
            self.head = n
            return
        curr = self.head
        while curr.next is not None:
            curr = curr.next
        curr.next = n

    def printList(self):
        curr = self.head
        while curr is not None:
            print curr.val
            curr = curr.next


l = LL()
l.insert(4)
l.insert(2)
l.insert(8)
l.printList()
isSubPath(l.head)
